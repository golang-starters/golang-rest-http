package main

import (
	"net/http"
	"os/user"

	"github.com/fabric8-services/fabric8-common/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/gzip"
	"github.com/golang-starters/golang-rest-http/app"
	"github.com/golang-starters/golang-rest-http/controller"
)

func main() {

	printUserInfo()

	// Create service
	service := goa.New("golang-foo")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(gzip.Middleware(9))
	service.Use(app.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount the status controller
	statusCtrl := controller.NewStatusController(service)
	app.MountStatusController(service, statusCtrl)

	http.Handle("/api/", service.Mux)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Start http server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error(nil, map[string]interface{}{
			"err": err,
		}, "unable to start server")
	}

}

func printUserInfo() {
	u, err := user.Current()
	if err != nil {
		log.Warn(nil, map[string]interface{}{
			"err": err,
		}, "failed to get current user")
	} else {
		log.Info(nil, map[string]interface{}{
			"username": u.Username,
			"uuid":     u.Uid,
		}, "Running as user name '%s' with UID %s.", u.Username, u.Uid)
		g, err := user.LookupGroupId(u.Gid)
		if err != nil {
			log.Warn(nil, map[string]interface{}{
				"err": err,
			}, "failed to lookup group")
		} else {
			log.Info(nil, map[string]interface{}{
				"groupname": g.Name,
				"gid":       g.Gid,
			}, "Running as as group '%s' with GID %s.", g.Name, g.Gid)
		}
	}

}
