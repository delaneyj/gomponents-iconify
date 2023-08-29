package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerStorageHardDriveOneDiskComputerDeviceElectronicsDiscDriveRaid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="7.5" rx="1"/><path d="m.59 8.08l1.74-6.8A1 1 0 0 1 3.3.5h7.4a1 1 0 0 1 1 .78l1.74 6.8M3.5 10.5h3"/><circle cx="10.5" cy="10.5" r=".5"/></g>`),
		g.Group(children),
	)
}