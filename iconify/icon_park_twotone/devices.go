package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDevices0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M23 43h20V5H14v10"/><path fill="#555" d="M5 15h18v28H5V15Z"/><path stroke-linecap="round" d="M13 37h2m13 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDevices0)"/>`),
		g.Group(children),
	)
}