package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        origin := ctx.Request.Header.Get("Origin")
        if origin == "http://localhost:3000/" || origin == "https://padang-v2-master-black.vercel.app/" {
            ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        } else {
            ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Fallback or default origin
        }

        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie, ngrok-skip-browser-warning")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

        if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(http.StatusNoContent)
            return
        }

        ctx.Next()
    }
}









// package middleware

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func CorsHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie, ngrok-skip-browser-warning")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
// 		if ctx.Request.Method == "OPTIONS" {
// 			ctx.AbortWithStatus(http.StatusNoContent)
// 		}
// 		ctx.Next()
// 	}
// }
