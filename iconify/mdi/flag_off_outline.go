package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3L5 6.89V21h2v-7h5.11l.61.61L13 16h1.11l6.73 6.73l1.27-1.27M7 12V8.89L10.11 12H7m2.2-6l-2-2H14l.4 2H20v10h-.8l-2-2h.8V8h-5.24l-.4-2H9.2Z"/>`),
		g.Group(children),
	)
}