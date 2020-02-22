build:
	@echo "  >  Building binary..."
	$$GOPATH/bin/qtdeploy -debug build  desktop 

linux-to-windows-build:
	@echo "  >  Building binary for windows_64_static..."
	$$GOPATH/bin/qtdeploy -docker -debug  build windows_64_static





      
