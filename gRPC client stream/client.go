import mediaGrpc /example.pd.go

func CustomerServiceToMail(ctx context.Context, name, email, subject, content string, files map[string][]byte) error {
	// 開啟gRPC資料流
	// 在gRPC產生的pb.go裡可以看到如何開啟stream
	// 在type ExampleServiceClient("%sClient", 你定的service名稱) interface
	// ExampleServiceClient.CustomerServiceToMail(ctx context.Context, opts ...grpc.CallOption) (ExampleService_CustomerServiceToMailClient, error)
	stream, err := mediaGrpc.ExampleServiceClient.CustomerServiceToMail(ctx)
	if err != nil {
		return err
	}

	// 可持續傳送資料，這裡傳送emailInfo
	err = stream.Send(&mediaGrpc.CustomerServiceToMailReq{Email: email, Subject: subject, Content: content, Name: name})
	if err != nil {
		return err
	}

	// 這裡傳送檔案
	for filename, file := range files {
		// 設定每次傳的data byte的size，這裡設32kb可自行更改
		chunksize := 32 << 10
		i := 0
		for ; len(file)/chunksize > i; i++ {
			// 每次傳送資料格式皆可任意設定，例如這裡只傳Data和Filename，不過一次最大只能傳4mb(gRPC預設，可更改)
			err = stream.Send(&mediaGrpc.CustomerServiceToMailReq{
				Data:     file[chunksize*i : chunksize*(i+1)],
				Filename: filename,
			})
			if err != nil {
				break
			}
		}
		if len(file[chunksize*i:]) != 0 {
			err = stream.Send(&mediaGrpc.CustomerServiceToMailReq{
				Data:     file[chunksize*i:],
				Filename: filename,
			})
		}
		if err != nil {
			return err
		}
	}

	// 結束後記得close stream，並且會由server端回傳(MediaEmpty, error)
	empty, err = stream.CloseAndRecv()
	return err
}