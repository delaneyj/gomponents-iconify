package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbStickTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.267 1.997a.75.75 0 0 1 .743.648l.007.102v4.386a2.252 2.252 0 0 1 1.48 1.946l.007.17v10.5a2.25 2.25 0 0 1-2.096 2.246l-.154.005h-6.5a2.25 2.25 0 0 1-2.245-2.096l-.005-.154V9.25a2.25 2.25 0 0 1 1.513-2.127V2.747a.75.75 0 0 1 .648-.743l.101-.007h6.501Zm-.75 1.5H9.516v3.502h5.001V3.497Z"/>`),
		g.Group(children),
	)
}