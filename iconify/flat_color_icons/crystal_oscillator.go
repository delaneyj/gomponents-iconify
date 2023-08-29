package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrystalOscillator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FF9800" d="M3 28h26v4H3zm0-12h26v4H3z"/><path fill="#2196F3" d="M43 11H20v26h23c1.1 0 2-.9 2-2V13c0-1.1-.9-2-2-2z"/><path fill="#64B5F6" d="M20 9h-2v30h2c1.1 0 2-.9 2-2V11c0-1.1-.9-2-2-2z"/>`),
		g.Group(children),
	)
}