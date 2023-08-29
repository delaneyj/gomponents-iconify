package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M288 192q0 28-14.5 51T235 278V106q24 12 38.5 35t14.5 51zM0 128h85L192 21v342L85 256H0V128z"/>`),
		g.Group(children),
	)
}