package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopScreenDeviceDisplayComputerDesktopElectronicsMonitorKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="8" x="1.5" y=".5" rx="1"/><path d="M12.28 11.05a1 1 0 0 0-.9-.55H2.62a1 1 0 0 0-.9.55l-.5 1a1 1 0 0 0 .05 1a1 1 0 0 0 .85.47h9.76a1 1 0 0 0 .85-.47a1 1 0 0 0 0-1ZM1.5 6h11"/></g>`),
		g.Group(children),
	)
}