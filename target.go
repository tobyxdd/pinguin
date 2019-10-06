package main

import "net"

func setTargetEnableState(appCtx *appContext, name string, enabled bool) {
	for i := range appCtx.AppConfig.Targets {
		if appCtx.AppConfig.Targets[i].Name == name {
			appCtx.AppConfig.Targets[i].Enabled = enabled
		}
	}
	if enabled {
		go func() {
			addr, err := net.ResolveIPAddr("ip", name)

			appCtx.InfoMapMutex.Lock()
			if err != nil {
				appCtx.InfoMap[name] = targetInfo{Error: err}
				return
			}
			appCtx.InfoMap[name] = targetInfo{Addr: addr}
			appCtx.InfoMapMutex.Unlock()

			_ = appCtx.Monitor.AddTarget(name, *addr)
		}()
	} else {
		appCtx.InfoMapMutex.Lock()
		delete(appCtx.InfoMap, name)
		appCtx.InfoMapMutex.Unlock()

		appCtx.Monitor.RemoveTarget(name)
	}
}

func addTarget(appCtx *appContext, name string) {
	for i := range appCtx.AppConfig.Targets {
		if appCtx.AppConfig.Targets[i].Name == name {
			return
		}
	}
	appCtx.AppConfig.Targets = append(appCtx.AppConfig.Targets, target{
		Enabled: true,
		Name:    name,
	})
	go func() {
		addr, err := net.ResolveIPAddr("ip", name)

		appCtx.InfoMapMutex.Lock()
		if err != nil {
			appCtx.InfoMap[name] = targetInfo{Error: err}
			return
		}
		appCtx.InfoMap[name] = targetInfo{Addr: addr}
		appCtx.InfoMapMutex.Unlock()

		_ = appCtx.Monitor.AddTarget(name, *addr)
	}()
}

func removeTarget(appCtx *appContext, name string) {
	for i := range appCtx.AppConfig.Targets {
		if appCtx.AppConfig.Targets[i].Name == name {
			appCtx.AppConfig.Targets = append(appCtx.AppConfig.Targets[:i], appCtx.AppConfig.Targets[i+1:]...)
			break
		}
	}
	appCtx.InfoMapMutex.Lock()
	delete(appCtx.InfoMap, name)
	appCtx.InfoMapMutex.Unlock()

	appCtx.Monitor.RemoveTarget(name)
}
