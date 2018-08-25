package handler

import (
	"github.com/gin-gonic/gin"
	"thresher/etcd/model"
)




func Get(g *gin.Context){

	 data := model.Arith2.Get(g,&model.Args{A:12,B:13},&model.Reply{})
	g.JSON(200, gin.H{
		"message": data.C,
	})

}



