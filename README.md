# gRPC-client-stream
使用gRPC client stream來傳送檔案的紀錄

client開啟的stream為interface共有Send和CloseAndRecv方法

server接收的stream為interface共有Recv和SendAndClose方法

靠這兩者便能完成gRPC client stream功能
