package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StyleGuideTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.819 17.259c.073.271.187.52.335.74h-.156c-1.104 0-2-.895-2.001-2l-.005-5.535L6.82 17.26ZM4 16.499c0 .152.014.301.04.446l-.056-.015a2.002 2.002 0 0 1-1.416-2.45l1.426-5.34L4 16.5Zm3.655.018a2.002 2.002 0 0 0 2.451 1.414l5.416-1.451a1.998 1.998 0 0 0 1.413-2.45L14.099 3.482a2.002 2.002 0 0 0-2.451-1.413l-5.416 1.45a1.998 1.998 0 0 0-1.413 2.449l2.836 10.55ZM9 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}