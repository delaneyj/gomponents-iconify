package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 272v150c0 9 7 15 15 15h84l169 129c27 28 47 18 47-20V147c0-39-20-47-47-19L99 257H15c-8 0-15 6-15 15z"/>`),
		g.Group(children),
	)
}