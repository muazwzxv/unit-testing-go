package app

import "unit_test/handlers"

func routes() {
	router.GET("/messages/:message_id", handlers.GetMessage)
	router.GET("/messages", handlers.GetAllMessages)
	router.POST("/messages", handlers.CreateMessage)
	router.PUT("/messages/:message_id", handlers.UpdateMessage)
	router.DELETE("/messages/:message_id", handlers.DeleteMessage)
}
