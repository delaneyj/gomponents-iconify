package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDevices0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M23 43h20V5H14v10"/><path fill="#fff" stroke="#fff" d="M5 15h18v28H5V15Z"/><path stroke="#000" stroke-linecap="round" d="M13 37h2"/><path stroke="#fff" stroke-linecap="round" d="M28 37h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDevices0)"/>`),
		g.Group(children),
	)
}