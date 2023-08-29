package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hulu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 12.8V22h-4.8v-8.1c0-.7-.6-1.3-1.3-1.3h-2.9c-.7 0-1.3.6-1.3 1.3V22H4.5V2h4.8v6.4c.3-.1.6-.2.9-.2H15c2.5 0 4.5 2.1 4.5 4.6Z"/>`),
		g.Group(children),
	)
}