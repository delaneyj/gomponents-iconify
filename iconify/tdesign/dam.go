package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .671l9 7.875V20h1v2h-9v-2h1V10h-4v10h1v2H2v-2h1V8.546l9-7.875Zm-7 19.33h3V10H5v10ZM6.662 8H17.34L12 3.328L6.661 8ZM19 10h-3v10h3V10Z"/>`),
		g.Group(children),
	)
}