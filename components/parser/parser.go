package parser

type URL []byte

type HTTPHost struct {
	Proto string
	Host  string
	Port  string
	Path  string
	Query string
}

func ParseURI(url URL) *HTTPHost {

	var proto string
	var host string
	var port string
	var path string
	var query string

	cur := 0

	//parse protocol
	for ; url[cur] != ':'; cur++ {
		proto += string(url[cur])
	}

	cur += 3

	//parse host
	for {
		if cur == len(url) || url[cur] == ':' || url[cur] == '/' {
			break
		}

		host += string(url[cur])

		cur++
	}

	//parse custom port
	if url[cur] == ':' {
		cur++
		for {
			if cur == len(url) || url[cur] == '/' {
				break
			}

			port += string(url[cur])

			cur++
		}
	}

	//parse if port is default
	if port == "" {
		if proto == "HTTP" {
			port = "80"
		}

		if proto == "HTTPS" {
			port = "443"
		}
		cur++
	}

	//parse http path
	for {
		if cur == len(url) || url[cur] == '?' {
			break
		}

		path += string(url[cur])
		cur++
	}

	//parse http query
	if url[cur] == '?' {
		cur++

		for ; cur != len(url); cur++ {
			query += string(url[cur])
		}
	}

	return &HTTPHost{
		Proto: proto,
		Host:  host,
		Port:  port,
		Path:  path,
		Query: query,
	}
}
