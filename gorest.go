package main

import (
	. "./services"
	//"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type Message struct {
	Msg string
	Err string
}

type SaveSpkdAddWSInput struct {
	USER_CODE              string
	BLPD_INDC              string
	CS_SPKD_PCN__CUST_NUMB string
	CS_SPKD_PCN__SUBR_NUMB string
	CS_SPKD_PCN__PACK_CODE string
	CS_SUBR_PCN__SUBR_TYPE string
	RD_TELP__TELP_TYPE     string
	SAVE_FLAG              string
}

type SaveSpkdAddWSoutput struct {
	CS_SPKD_PCN__PACK_CODE       string
	CS_PKPL_PCN__PACK_DESC       string
	CS_PACK_TYPE__PACK_TYPE_DESC string
	CS_SPKD_PCN__PACK_STRT_DTTM  string
	CS_SPKD_PCN__PACK_END_DTTM   string
	CS_SPKD_PCN__DISC_CODE       string
	TBL_OCCR                     string
}

func main() {
	//num_cores := 8
	//for i := 0; i < num_cores; i++ {
	//	go func() {
	handler := rest.ResourceHandler{
		PreRoutingMiddlewares: []rest.Middleware{
			&rest.CorsMiddleware{
				RejectNonCorsRequests: false,
				OriginValidator: func(origin string, request *rest.Request) bool {
					return origin == "http://10.35.40.41"
				},
				AllowedMethods:                []string{"GET", "POST", "PUT"},
				AllowedHeaders:                []string{"Accept", "Content-Type", "X-Custom-Header"},
				AccessControlAllowCredentials: true,
				AccessControlMaxAge:           3600,
			},
		},
	}
	handler.SetRoutes(
		&rest.Route{"POST", "/post", SaveSpkdAddWSService},
		&rest.Route{"OPTIONS", "/post", SaveSpkdAddWSService},
	)
	http.ListenAndServe(":8080", &handler)
	//}()
	//}
}

func SaveSpkdAddWSService(w rest.ResponseWriter, r *rest.Request) {

	input := SaveSpkdAddWSInput{}
	err := r.DecodeJsonPayload(&input)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		println(err.Error())
		return
	}
	if input.USER_CODE == "" {
		rest.Error(w, "USER_CODE required", 400)
		println("USER_CODE required")
		return
	}
	if input.BLPD_INDC == "" {
		rest.Error(w, "BLPD_INDC required", 400)
		println("BLPD_INDC required")
		return
	}

	var saveSpkdAddWS = SaveSpkdAddWS{USER_CODE: input.USER_CODE, BLPD_INDC: input.BLPD_INDC, CS_SPKD_PCN__CUST_NUMB: input.CS_SPKD_PCN__CUST_NUMB, CS_SPKD_PCN__SUBR_NUMB: input.CS_SPKD_PCN__SUBR_NUMB, CS_SPKD_PCN__PACK_CODE: input.CS_SPKD_PCN__PACK_CODE, RD_TELP__TELP_TYPE: input.RD_TELP__TELP_TYPE, SAVE_FLAG: input.SAVE_FLAG}

	var bodyElement *SaveSpkdAddWSResponse
	var faultElement *FaultResponse
	var tuxerr error
	var state string

	bodyElement, faultElement, tuxerr, state = CallTux(saveSpkdAddWS)

	if tuxerr != nil {
		println(tuxerr.Error())
		w.WriteJson(&SaveSpkdAddWSoutput{
			CS_SPKD_PCN__PACK_CODE:       "fault",
			CS_PKPL_PCN__PACK_DESC:       "fault",
			CS_PACK_TYPE__PACK_TYPE_DESC: "fault",
			CS_SPKD_PCN__PACK_STRT_DTTM:  "fault",
			CS_SPKD_PCN__PACK_END_DTTM:   "fault",
			CS_SPKD_PCN__DISC_CODE:       "fault",
			TBL_OCCR:                     "0",
		})
	}
	if faultElement != nil {
		println("fault error")
	}
	if state == "body" {
		println("body state")
	} else {
		println(state)
	}
	if bodyElement == nil {
		println("no body")
	} else {
		w.WriteJson(&SaveSpkdAddWSoutput{
			CS_SPKD_PCN__PACK_CODE:       bodyElement.CS_SPKD_PCN__PACK_CODE,
			CS_PKPL_PCN__PACK_DESC:       bodyElement.CS_PKPL_PCN__PACK_DESC,
			CS_PACK_TYPE__PACK_TYPE_DESC: bodyElement.CS_PACK_TYPE__PACK_TYPE_DESC,
			CS_SPKD_PCN__PACK_STRT_DTTM:  bodyElement.CS_SPKD_PCN__PACK_STRT_DTTM,
			CS_SPKD_PCN__PACK_END_DTTM:   bodyElement.CS_SPKD_PCN__PACK_END_DTTM,
			CS_SPKD_PCN__DISC_CODE:       bodyElement.CS_SPKD_PCN__DISC_CODE,
			TBL_OCCR:                     bodyElement.TBL_OCCR,
		})
	}

}

//func PostParameter(w rest.ResponseWriter, r *rest.Request) {
//	fmt.Printf("HEADER: %s\n", r.Header)
//	input := SaveSpkdAddWSInput{}
//	err := r.DecodeJsonPayload(&input)

//	if err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		println(err.Error())
//		return
//	}
//	if input.USER_CODE == "" {
//		rest.Error(w, "USER_CODE required", 400)
//		println("USER_CODE required")
//		return
//	}
//	if input.BLPD_INDC == "" {
//		rest.Error(w, "BLPD_INDC required", 400)
//		println("BLPD_INDC required")
//		return
//	}

//	w.WriteJson(&SaveSpkdAddWSoutput{
//		CS_SPKD_PCN__PACK_CODE:       "testa",
//		CS_PKPL_PCN__PACK_DESC:       "testb",
//		CS_PACK_TYPE__PACK_TYPE_DESC: "testc",
//		CS_SPKD_PCN__PACK_STRT_DTTM:  "testd",
//		CS_SPKD_PCN__PACK_END_DTTM:   "teste",
//		CS_SPKD_PCN__DISC_CODE:       "testf",
//		TBL_OCCR:                     "999",
//	})
//}
