package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FountainPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#94989b" d="m12.3 43.8l7.9 7.9l-6.1 10.1S7.7 61.1 1.6 64l10-10c.5 0 .9-.1 1.3-.5c.7-.7.7-1.7 0-2.4s-1.7-.7-2.4 0c-.4.4-.5.9-.5 1.3l-10 10c2.9-6.2 2.2-12.5 2.2-12.5l10.1-6.1"/><path fill="#3e4347" d="M61.5 2.5c3.3 3.3 3.3 8.6 0 11.9L56.9 19S55.1 24 46 33.1C36.8 42.2 31.8 44 31.8 44l-9.7 9.7l-11.8-11.9l9.7-9.7s1.8-5 10.9-14.1S45 7.1 45 7.1l4.7-4.7c3.3-3.2 8.6-3.2 11.8.1z"/><path fill="#94989b" d="m45.008 7.149l1.555-1.556l11.88 11.88l-1.556 1.555zM18.404 33.752l1.554-1.557l11.89 11.87l-1.553 1.556z"/>`),
		g.Group(children),
	)
}