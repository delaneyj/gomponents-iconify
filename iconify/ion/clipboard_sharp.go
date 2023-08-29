package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M420 48h-68V28a12 12 0 0 0-12-12H172a12 12 0 0 0-12 12v20H92a12 12 0 0 0-12 12v424a12 12 0 0 0 12 12h328a12 12 0 0 0 12-12V60a12 12 0 0 0-12-12Zm-84.13 64H176.13V80h159.74Z"/>`),
		g.Group(children),
	)
}