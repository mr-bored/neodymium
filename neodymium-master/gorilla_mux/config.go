package gorilla_mux

func GetConfig () configStruct {

	const HTTP_SERVER_IP = "localhost";
	const HTTP_SERVER_PORT = 4492;
	const HTTP_SSL_ACTIVE = false;
	return configStruct{ HTTP_SERVER_IP, HTTP_SERVER_PORT, HTTP_SSL_ACTIVE };
}

type ServerconfigStruct struct{
	ip string `ip`
	port int8 `port`
	ssl_active bool `ssl_active`
}

