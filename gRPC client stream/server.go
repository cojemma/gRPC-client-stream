import mediaGrpc /example.pd.go

// 在gRPC產生的pb.go裡的格式
// 在type ExampleServiceServer("%sServer", 你定的service名稱) interface
// CustomerServiceToMail(ExampleService_CustomerServiceToMailServer) error
func CustomerServiceToMail(stream mediaGrpc.ExampleService_CustomerServiceToMailServer) (err error) {
	files := make(map[string][]byte)
	emailInfo := struct {
		Subject string
		ToEmail []string
		Content string
	}
	// 獲取stream的context
	ctx := stream.Context()

	// 透過無限迴圈，來持續的從stream資料流Recv資料
	for {
		// req為client傳來的CustomerServiceToMailReq的struct
		req, err := stream.Recv()
		if err != nil {
			// 當所有資料街傳送完畢，則err == io.EOF，並跳出迴圈
			if err == io.EOF {
				break
			}
			break
		}
		// 判定接收的資料為emailInfo
		if req.Subject != "" {
			emailInfo.Subject = req.Subject
			emailInfo.ToEmail = []string{}
			content := helpers.EmailHead + req.Content + "\n\r\n\r\n" + "回覆 Email: " + req.Email + "\r\n會員名稱 : " + req.Name + helpers.EmailTail
			emailInfo.Content = content
		}
		// 判定接收的資料為檔案則寫入files map
		if req.Filename != "" {
			filename := req.Filename
			if _, ok := files[filename]; !ok {
				files[filename] = []byte{}
			}
			files[filename] = append(files[filename], req.Data...)
		}
	}

	// 最後server端記得關閉stream並回傳MediaEmpty
	defer stream.SendAndClose(&mediaGrpc.MediaEmpty{})
	return nil
}