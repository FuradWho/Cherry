package Makefile

/*
	虽然 Makefile 是一个很老的项目管理工具，但它仍然是最优秀的项目管理工具。
	所以，一个 Go 项目在其根目录下应该有一个 Makefile 工具，用来对项目进行管理，Makefile 通常用来执行静态代码检查、单元测试、编译等功能。

	我还有一条建议：直接执行 make 时，执行如下各项 format -> lint -> test -> build，
	如果是有代码生成的操作，还可能需要首先生成代码 gen -> format -> lint -> test -> build。
	在实际开发中，我们可以将一些重复性的工作自动化，并添加到 Makefile 文件中统一管理。

*/
