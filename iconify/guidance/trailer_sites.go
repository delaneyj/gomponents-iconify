package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrailerSites(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18.5 18.5H24m-5.5 0v-14h-15S.5 9 .5 16.616V18.5H7m11.5 0H12m9 2h3m-17-2a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m-2-11v4m-6 0a23.99 23.99 0 0 1 1.172-4H15.5v4H4Z"/>`),
		g.Group(children),
	)
}