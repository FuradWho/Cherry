package configs

/*
	与部署阶段相关的目录，这些目录可以存放部署相关的文件。

	/configs这个目录用来配置文件模板或默认配置。
	例如，可以在这里存放 confd 或 consul-template 模板文件。这里有一点要注意，配置中不能携带敏感信息，这些敏感信息，我们可以用占位符来替代，
	例如：apiVersion: v1 user: username: ${CONFIG_USER_USERNAME} # iam 用户名 password: ${CONFIG_USER_PASSWORD} # iam 密码

*/
