package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v20.063L8.219 17.28L6.78 18.72l8.5 8.5l.719.687l.719-.687l8.5-8.5l-1.438-1.438L17 24.063V4z"/>`),
		g.Group(children),
	)
}