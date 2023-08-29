package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbEngOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><path fill="#ce1124" d="M215 0h82v512h-82z"/><path fill="#ce1124" d="M0 215h512v82H0z"/>`),
		g.Group(children),
	)
}