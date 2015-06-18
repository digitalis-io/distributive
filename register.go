package main

func registerCheck(name string, work Worker, numParams int) {
	workers[name] = work
	parameterLength[name] = numParams
}

func registerChecks() {
	registerCheck("command", command, 1)
	registerCheck("running", running, 1)
	registerCheck("phpconfig", phpConfig, 2)
	registerCheck("file", file, 1)
	registerCheck("directory", directory, 1)
	registerCheck("symlink", symlink, 1)
	registerCheck("checksum", checksum, 3)
	registerCheck("permissions", permissions, 2)
	registerCheck("filecontains", fileContains, 2)
	registerCheck("diskusage", diskUsage, 2)
	registerCheck("memoryusage", memoryUsage, 1)
	registerCheck("swapusage", swapUsage, 1)
	registerCheck("temp", temp, 1)
	registerCheck("port", port, 1)
	registerCheck("interface", interfaceExists, 1)
	registerCheck("up", up, 1)
	registerCheck("ip4", ip4, 2)
	registerCheck("ip6", ip6, 2)
	registerCheck("gateway", gateway, 1)
	registerCheck("gatewayinterface", gatewayInterface, 1)
	registerCheck("host", host, 1)
	registerCheck("tcp", tcp, 1)
	registerCheck("udp", udp, 1)
	registerCheck("tcptimeout", tcpTimeout, 2)
	registerCheck("udptimeout", udpTimeout, 2)
	registerCheck("routingtabledestination", routingTableDestination, 1)
	registerCheck("routingtableinterface", routingTableInterface, 1)
	registerCheck("routingtablegateway", routingTableGateway, 1)
	registerCheck("module", module, 1)
	registerCheck("kernelparameter", kernelParameter, 1)
	registerCheck("dockerimage", dockerImage, 1)
	registerCheck("dockerrunning", dockerRunning, 1)
	registerCheck("groupexists", groupExists, 1)
	registerCheck("useringroup", userInGroup, 2)
	registerCheck("groupid", groupId, 2)
	registerCheck("userexists", userExists, 1)
	registerCheck("userhasuid", userHasUID, 2)
	registerCheck("userhasgid", userHasGID, 2)
	registerCheck("userhasusername", userHasUsername, 2)
	registerCheck("userhasname", userHasName, 2)
	registerCheck("userhashomedir", userHasHomeDir, 2)
	registerCheck("installed", installed, 1)
	registerCheck("repoexists", repoExists, 2)
	registerCheck("repoexistsuri", repoExistsURI, 2)
	registerCheck("pacmanignore", pacmanIgnore, 1)
	registerCheck("systemctlloaded", systemctlLoaded, 1)
	registerCheck("systemctlactive", systemctlActive, 1)
	registerCheck("systemctlsockpath", systemctlSockPath, 1)
	registerCheck("systemctlsockunit", systemctlSockUnit, 1)
	registerCheck("systemctltimer", systemctlTimer, 1)
	registerCheck("systemctltimerloaded", systemctlTimerLoaded, 1)
	registerCheck("systemctlunitfilestatus", systemctlUnitFileStatus, 2)
}
