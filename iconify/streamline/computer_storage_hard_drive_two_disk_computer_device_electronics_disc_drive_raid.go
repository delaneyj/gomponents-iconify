package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerStorageHardDriveTwoDiskComputerDeviceElectronicsDiscDriveRaid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 13.5a3 3 0 0 0 0-6h-7a3 3 0 0 0 0 6Z"/><path d="M.58 9.79L2.17 2.1a2 2 0 0 1 2-1.6h5.7a2 2 0 0 1 2 1.6l1.59 7.69m-9.96.71h3"/><circle cx="10.5" cy="10.5" r=".5"/></g>`),
		g.Group(children),
	)
}