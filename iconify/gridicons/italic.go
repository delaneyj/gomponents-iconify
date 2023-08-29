package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.536 5l-.427 2h1.5L9.263 18h-1.5l-.427 2h6.128l.426-2h-1.5l2.347-11h1.5l.427-2z"/>`),
		g.Group(children),
	)
}