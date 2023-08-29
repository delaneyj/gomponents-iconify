package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestartCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="12" r="10" opacity=".5"/><path d="M16.402 6.286a.75.75 0 0 0-1.287-.524l-.75.768a5.666 5.666 0 0 0-6.437 1.208c-2.237 2.29-2.237 5.996 0 8.287a5.669 5.669 0 0 0 8.144 0a5.934 5.934 0 0 0 1.634-4.874a.75.75 0 0 0-1.49.182c.16 1.3-.25 2.653-1.217 3.643a4.169 4.169 0 0 1-5.998 0c-1.668-1.707-1.668-4.483 0-6.19a4.17 4.17 0 0 1 4.224-1.089l-.761.78A.75.75 0 0 0 13 9.75h2.652a.75.75 0 0 0 .75-.75V6.286Z"/></g>`),
		g.Group(children),
	)
}