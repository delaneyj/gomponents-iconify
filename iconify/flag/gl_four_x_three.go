package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h640v480H0z"/><path fill="#d00c33" d="M0 240h640v240H0zm80 0a160 160 0 1 0 320 0a160 160 0 0 0-320 0"/>`),
		g.Group(children),
	)
}