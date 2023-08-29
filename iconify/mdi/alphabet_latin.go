package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabetLatin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6v12h2v-.69c.37.42.9.69 1.5.69h.5c1.65 0 3-1.35 3-3v-3c0-1.64-1.35-3-3-3h-.5c-.6 0-1.13.27-1.5.7V6M5 9v2h3c.57 0 1 .43 1 1H7c-1.64 0-3 1.36-3 3c0 1.65 1.36 3 3 3h4v-6c0-1.64-1.35-3-3-3m8 2h1c.57 0 1 .43 1 1v3c0 .57-.43 1-1 1h-1c-.57 0-1-.43-1-1v-3c0-.57.43-1 1-1m-9 3h2v2H7c-.57 0-1-.43-1-1c0-.57.43-1 1-1Z"/>`),
		g.Group(children),
	)
}