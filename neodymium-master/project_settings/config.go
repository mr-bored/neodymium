package project_settings

func GetProjectConfig () configStruct {

	const PROJECT_VERSION = 1.0;
	const AUTOR = "sjcm_98";

	return configStruct{ PROJECT_VERSION, AUTOR };
}

type configStruct struct{
	project_version float32 `ip`
	autor string `port`
}

