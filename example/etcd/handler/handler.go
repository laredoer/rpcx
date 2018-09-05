package handler

import (
	"github.com/gin-gonic/gin"
	"thresher/etcd/model"
)

func Get(g *gin.Context){
	data,err := model.Arith2.Get(g,&model.Args{A:12,B:13},&model.Reply{})
	if err != nil {
		g.JSON(500, gin.H{
			"message": err,
			"data": data,
		})
	}
	g.JSON(200,gin.H{
		"message": err,
		"data":data,
	})

}



