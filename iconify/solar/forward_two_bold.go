package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 6.25a.75.75 0 0 1 .75.75c0 1.756.715 2.78 1.563 3.391c.887.639 1.974.859 2.687.859h8.19l-3.72-3.72a.75.75 0 0 1 1.06-1.06l5 5a.75.75 0 0 1 0 1.06l-5 5a.75.75 0 1 1-1.06-1.06l3.72-3.72H9.5c-.953 0-2.367-.28-3.563-1.141C4.702 10.719 3.75 9.244 3.75 7a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}