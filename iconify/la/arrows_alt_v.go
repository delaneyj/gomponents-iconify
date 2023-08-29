package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsAltV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.094l-.719.687l-8 8L8.72 12.22L15 5.938v20.125L8.719 19.78L7.28 21.22l8 8l.719.687l.719-.687l8-8l-1.438-1.438L17 26.063V5.938l6.281 6.28l1.438-1.437l-8-8z"/>`),
		g.Group(children),
	)
}