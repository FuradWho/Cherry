package pkg

/*

	/internal/pkg：存放项目内可共享，项目外不共享的包。
		这些包提供了比较基础、通用的功能，例如工具、错误码、用户验证等功能。
	我的建议是，一开始将所有的共享代码存放在 /internal/pkg 目录下，当该共享代码做好了对外开发的准备后，再转存到/pkg目录下。

	/internal/pkg 目录存放项目内可共享的包，通常可以包含如下目录：
		/internal/pkg/code：项目业务 Code 码。
		/internal/pkg/validation：一些通用的验证函数。
		/internal/pkg/middleware：HTTP 处理链。
*/
