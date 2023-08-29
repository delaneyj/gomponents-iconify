package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbEngFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h640v480H0z"/><path fill="#ce1124" d="M281.6 0h76.8v480h-76.8z"/><path fill="#ce1124" d="M0 201.6h640v76.8H0z"/>`),
		g.Group(children),
	)
}