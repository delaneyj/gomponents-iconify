package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdeogramCjk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4v2H4v4h2V8h12v2h2V6h-7V4m-5 6v2h5.59l-2 2H4v2h7v2h-1v2h3v-4h7v-2h-5.79L16 12.21V10Z"/>`),
		g.Group(children),
	)
}