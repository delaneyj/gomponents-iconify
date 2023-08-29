package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.483 12.482h3.034L12 8.831z"/><path fill="currentColor" d="M12 3.074L3.689 6.038l1.268 10.987l7.043 3.9l7.043-3.9l1.268-10.987L12 3.074zm5.187 13.621H15.25l-1.045-2.606h-4.41L8.75 16.695H6.813L12 5.047l5.187 11.648z"/>`),
		g.Group(children),
	)
}