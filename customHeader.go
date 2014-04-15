package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
)

type MyCorsMiddleware struct{}

func (mw *MyCorsMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
    return func(writer rest.ResponseWriter, request *rest.Request) {

        corsInfo := request.GetCorsInfo()
            // check the request methods
            allowedMethods := map[string]bool{
                "GET":  true,
                "POST": true,
                "PUT":  true,
                // don't allow DELETE, for instance
            }

            // check the request headers
            allowedHeaders := map[string]bool{
                "Accept":          true,
                "Content-Type":    true,
                "X-Custom-Header": true,
            }
            for _, requestedHeader := range corsInfo.AccessControlRequestHeaders {
                if !allowedHeaders[requestedHeader] {
                    rest.Error(writer, "Invalid Preflight Request", http.StatusForbidden)
                    return
                }
            }

            for allowedMethod, _ := range allowedMethods {
                writer.Header().Add("Access-Control-Allow-Methods", allowedMethod)
            }
            for allowedHeader, _ := range allowedHeaders {
                writer.Header().Add("Access-Control-Allow-Headers", allowedHeader)
            }
            writer.Header().Set("Access-Control-Allow-Origin", corsInfo.Origin)
            writer.Header().Set("Access-Control-Allow-Credentials", "true")
            writer.Header().Set("Access-Control-Max-Age", "3600")
            writer.WriteHeader(http.StatusOK)
            return

    }
}

func main() {

    handler := rest.ResourceHandler{
        PreRoutingMiddlewares: []rest.Middleware{
            &MyCorsMiddleware{},
        },
    }
    handler.SetRoutes(
        &rest.Route{"GET", "/countries", GetAllCountries},
    )
    http.ListenAndServe(":8080", &handler)
}

type Country struct {
    Code string
    Name string
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(
        []Country{
            Country{
                Code: "FR",
                Name: "France",
            },
            Country{
                Code: "US",
                Name: "United States",
            },
        },
    )
}