package routers

import (
	"backend/src/controllers"
	"net/http"
)

var routerUsers = []RouterApi{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		HandlerFunction:        controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		HandlerFunction:        controllers.FindAllUsers,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodGet,
		HandlerFunction:        controllers.FindUserById,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodPut,
		HandlerFunction:        controllers.UpdateUserById,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		HandlerFunction:        controllers.DeleteUserById,
		RequiresAuthentication: false,
	},
}
