package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M243 47q-15 15-2 30l77 94H21q-21 0-21 21t21 21h297l-77 94q-6 7-5 16t7 14q7 6 16 5t14-7l111-143L273 51q-13-16-30-4z"/>`),
		g.Group(children),
	)
}