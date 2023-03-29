userspace.myRpc.dir := userspace/myRpc

# remote procedure interface between client and server
userspace.myRpc.types        = $(patsubst %, $(userspace.myRpc.dir)/%, myRpc.t)
userspace.myRpc.packgage = $(userspace.myRpc.dir)/myRpc.go


# client build definitions
userspace.bankingclient.go     = bankingclient.go
userspace.bankingclient.src    = $(patsubst %, $(userspace.myRpc.dir)/%, $(userspace.bankingclient.go))
userspace.bankingclient.target = $(UDIR)/$(userspace.myRpc.dir)/bankingclient


# server build definitions
userspace.bankingserver.go    = bankingserver.go
userspace.bankingserver.src   = $(patsubst %, $(userspace.myRpc.dir)/%, $(userspace.bankingserver.go))
userspace.bankingserver.target= $(UDIR)/$(userspace.myRpc.dir)/bankingserver

# install these executables
INSTALL_ETHOSROOT_PROGRAMS        += $(userspace.bankingclient.target)
INSTALL_ETHOSROOT_SYSTEM_PROGRAMS += $(userspace.bankingserver.target)

# compile the types
$(userspace.myRpc.dir)/myRpc.go: $(userspace.myRpc.types)
	$(ETN2GO) $(@D) myRpc $^

$(userspace.myRpc.package) : $(userspace.myRpc.dir)/myRpc.go


# compile the client
$(userspace.bankingclient.target): $(userspace.bankingclient.src) $(userspace.myRpc.package)
	sh bin/ethosGo  $(userspace.bankingclient.src)

# compile the server
$(userspace.bankingserver.target): $(userspace.bankingserver.src) $(userspace.myRpc.package)
	sh bin/ethosGo $(userspace.bankingserver.src)


# build everything
userspace.myRpc.all: $(userspace.bankingclient.target) $(userspace.bankingserver.target)

# install types, service,
userspace.myRpc.installfs:
	bin/ethosTypeInstall $(userspace.myRpc.dir)/myRpc
	bin/ethosServiceInstall myRpc MyRpc
	ethosStringEncode /system/programs/bankingserver    > $(install.ethosRoot.init.services)/bankingserver

# remove build artifacts
userspace.myRpc.clean:
	rm -rf $(userspace.myRpc.dir)/myRpc
	rm -f  $(userspace.myRpc.dir)/myRpc.go

