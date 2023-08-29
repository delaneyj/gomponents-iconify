package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardWirelessRemoteDeviceComputerWirelessElectronicsQwertyKeyboardBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="3" rx="1"/><path d="M3.5 8.25h7m-7-2.5h7"/></g>`),
		g.Group(children),
	)
}