package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageHaskell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.08 19l4.67-7l-4.67-7h3.5l4.67 7l-4.67 7h-3.5m4.67 0l4.67-7l-4.67-7h3.5l9.34 14h-3.5l-2.92-4.37L10.25 19h-3.5m11.28-4.08l-1.53-2.34h5.42v2.34h-3.89m-2.33-3.5l-1.56-2.34h7.78v2.34H15.7Z"/>`),
		g.Group(children),
	)
}