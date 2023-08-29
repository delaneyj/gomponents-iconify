package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerStorageHardDiskDeviceDiscDriveComputerDiskElectronicsPlatterTurntableRaid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="m4.5 10.5l1.75-2.75M4.09 7A2.93 2.93 0 0 1 4 5.16a3 3 0 1 1 4.67 3.27M7.5 11H10"/></g>`),
		g.Group(children),
	)
}