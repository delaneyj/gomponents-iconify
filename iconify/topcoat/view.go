package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func View(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M15.3 20.1c0 3.1 2.6 5.7 5.7 5.7s5.7-2.6 5.7-5.7s-2.6-5.7-5.7-5.7s-5.7 2.6-5.7 5.7zm8.1 12.3C30.1 30.9 40.5 22 40.5 22s-7.7-12-18-13.3c-.6-.1-2.6-.1-3-.1c-10 1-18 13.7-18 13.7s8.7 8.6 17 9.9c.9.4 3.9.4 4.9.2zM11.1 20.7c0-5.2 4.4-9.4 9.9-9.4s9.9 4.2 9.9 9.4S26.5 30 21 30s-9.9-4.2-9.9-9.3z"/>`),
		g.Group(children),
	)
}