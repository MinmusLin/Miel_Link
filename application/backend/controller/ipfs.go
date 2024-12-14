package controller

import (
    "backend/pkg"
    "fmt"

    "github.com/gin-gonic/gin"
)

func IpfsUpload(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Upload failed: " + err.Error(),
        })
        return
    }
    err = c.SaveUploadedFile(file, fmt.Sprintf("./files/uploadfiles/%v", file.Filename))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Upload failed: " + err.Error(),
        })
        return
    }
    cid, err := pkg.IpfsAdd(file.Filename)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Upload failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Upload successfully",
        "cid":     cid,
    })
}

func IpfsDownload(c *gin.Context) {
    cid := c.PostForm("cid")
    filename := c.PostForm("filename")
    err := pkg.IpfsGet(cid, filename)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Download failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Download successfully",
        "data":    filename,
    })
}

func Download(c *gin.Context) {
    filename := c.Query("filename")
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
    c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
    c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", filename))
    c.File("./files/downloadfiles/" + filename)
}
