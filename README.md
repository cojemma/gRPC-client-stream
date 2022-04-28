# gRPC-client-stream
使用gRPC client stream來傳送檔案的紀錄

client開啟的stream為interface共有
type MediaService_CustomerServiceToMailClient interface {
	Send(*CustomerServiceToMailReq) error
	CloseAndRecv() (*MediaEmpty, error)
	grpc.ClientStream
}

server接收的stream為
type MediaService_CustomerServiceToMailServer interface {
	SendAndClose(*MediaEmpty) error
	Recv() (*CustomerServiceToMailReq, error)
	grpc.ServerStream
}
