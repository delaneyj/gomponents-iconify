package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21.05l-9-7l1.65-1.25L12 18.5l7.35-5.7L21 14.05l-9 7ZM12 16L3 9l9-7l9 7l-9 7Zm0-7Zm0 4.45L17.75 9L12 4.55L6.25 9L12 13.45Z"/>`),
		g.Group(children),
	)
}