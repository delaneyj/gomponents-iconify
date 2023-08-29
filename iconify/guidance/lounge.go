package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lounge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M22 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-9L8.621 8.985a2 2 0 0 1 1.94-2.485h.939l4 8m1 5.5h-5a4.643 4.643 0 0 1-4.504-3.517L3 .5m6.195 4s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 8.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}