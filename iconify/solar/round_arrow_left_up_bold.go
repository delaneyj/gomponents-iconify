package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundArrowLeftUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm2.25-13a.75.75 0 0 0-.75-.75H9a.75.75 0 0 0-.75.75v4.5a.75.75 0 0 0 1.5 0v-2.69l4.72 4.72a.75.75 0 1 0 1.06-1.06l-4.72-4.72h2.69a.75.75 0 0 0 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}