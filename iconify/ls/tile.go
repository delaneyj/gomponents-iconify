package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 311h297V14H0v297zm369 0h297V14H369v297zM0 680h297V383H0v297zm369 0h297V383H369v297z"/>`),
		g.Group(children),
	)
}