package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelBarrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2h-2.3l-.16.07l-.17.11a.8.8 0 0 0-.13.13a.86.86 0 0 0-.1.16a.71.71 0 0 0-.08.18v.09L17.38 6h-1.14l-3.12-3.11a3.06 3.06 0 0 0-4.24 0L5.76 6H3a1 1 0 0 0-.87.5a1 1 0 0 0 0 1l4 7l-.77 1.5A2.2 2.2 0 0 0 5 16a3 3 0 1 0 3 3a3 3 0 0 0-.85-2.08l1-2l1.38-.14l3.94 5.91A2.93 2.93 0 0 0 16 22a3.18 3.18 0 0 0 1.13-.21a3 3 0 0 0 1.87-3.3L18 13l1.79-9H21a1 1 0 0 0 0-2ZM5 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm5.3-15.71a1 1 0 0 1 1.4 0L13.41 6H8.59Zm-2.75 8.65L4.72 8H17l-.82 4.08Zm8.81 7a1 1 0 0 1-1.2-.38l-3.34-5l4.37-.43l.81 4.69a1 1 0 0 1-.64 1.11Z"/>`),
		g.Group(children),
	)
}