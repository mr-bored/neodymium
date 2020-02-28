package mongodb

func GetConfig () configStruct {

	const MONGODB_IP = "localhost";
	const MONGODB_PORT = 4492;

	return configStruct{ MONGODB_IP, MONGODB_PORT };
}

type configStruct struct{
	ip string `ip`
	port int8 `port`
}

