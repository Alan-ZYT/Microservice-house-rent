package models

import (
	"github.com/weilaihui/fdfs_client"
	"github.com/astaxie/beego"
)

//通过文件名的方式进行上传
func UploadByFilename( filename string)(GroupName,RemoteFileId string ,err error ) {
	//通过配置文件创建fdfs操作句柄
	fdfsClient, thiserr :=fdfs_client.NewFdfsClient("./conf/client.conf")
	if thiserr  !=nil{
		//说一下那里出问题了
		beego.Info("UploadByFilename( ) fdfs_client.NewFdfsClient  err",err)
		GroupName = ""
		RemoteFileId = ""
		err = thiserr
		return
	}

	//unc (this *FdfsClient) UploadByFilename(filename string) (*UploadFileResponse, error)
	//通过句柄上传文件（被上传的文件）

	uploadResponse, thiserr := fdfsClient.UploadByFilename(filename)
	if thiserr !=nil{
		beego.Info("UploadByFilename( ) fdfsClient.UploadByFilename(filename)  err",err)
		GroupName = ""
		RemoteFileId = ""
		err = thiserr
		return
	}

	beego.Info(uploadResponse.GroupName)
	beego.Info(uploadResponse.RemoteFileId)
	//回传
	return  uploadResponse.GroupName , uploadResponse.RemoteFileId ,nil

}




//功能函数 操作fdfs上传二进制文件
func UploadByBuffer(filebuffer []byte, fileExtName string)(GroupName,RemoteFileId string ,err error ){

	//通过配置文件创建fdfs操作句柄
	fdfsClient, thiserr :=fdfs_client.NewFdfsClient("./conf/client.conf")
	if thiserr  !=nil{
		beego.Info("UploadByBuffer( ) fdfs_client.NewFdfsClient  err",err)
		GroupName = ""
		RemoteFileId = ""
		err = thiserr
		return
	}

	//通过句柄上传二进制的文件
	uploadResponse, thiserr :=fdfsClient.UploadByBuffer(filebuffer,fileExtName)
	if thiserr  !=nil{
		beego.Info("UploadByBuffer( ) fdfs_client.UploadByBuffer  err",err)
		GroupName = ""
		RemoteFileId = ""
		err = thiserr
		return
	}
	beego.Info(uploadResponse.GroupName)
	beego.Info(uploadResponse.RemoteFileId)
	//回传入
	return uploadResponse.GroupName,uploadResponse.RemoteFileId,nil

}
