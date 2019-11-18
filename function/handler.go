package function

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	html := `
	<html>
<head>
<title>OpenFaaS  + Okteto in San Diego</title>
</head>
<body style="background-color: gray;">
  <h1 style="color: white;"> Welcome to San Diego!</h1>
  <img src="https://live.staticflickr.com/4090/5040304297_eb5743f76a_b.jpg" />
  <br />
  
  <!--
  <span style="margin: 20px;">
      <img src="https://blog.alexellis.io/content/images/2017/08/faas_side.png" width="400px" />
  </span>
  <span style="margin: 20px;">
      <img src="https://okteto.com/okteto-logo-light-h-1.1.png" width="400px" />  
  </span>
  -->
</body>
</html>
	`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
