package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentSaveDiskFloppyElectronicsDeviceDiscComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 12.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-9l3-3h9a1 1 0 0 1 1 1Z"/><path d="M3.5 8.5h7v5h-7zm1-8h6v4h-6z"/></g>`),
		g.Group(children),
	)
}