package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageFortran(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4v2h1c.55 0 1 .45 1 1v10c0 .55-.45 1-1 1H5v2h9v-2h-2c-.55 0-1-.45-1-1v-4h2c.55 0 1 .45 1 1v2h2V8h-2v2c0 .55-.45 1-1 1h-2V6h5c1.11 0 2 1.34 2 3v1h2V4Z"/>`),
		g.Group(children),
	)
}