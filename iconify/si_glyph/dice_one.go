package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.763 0H3.178c-1.18 0-2.135.966-2.135 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 14.763 0zM9.017 11.155c-1.696 0-3.069-1.406-3.069-3.14c0-1.734 1.373-3.14 3.069-3.14c1.694 0 3.069 1.406 3.069 3.14c0 1.734-1.375 3.14-3.069 3.14z"/>`),
		g.Group(children),
	)
}