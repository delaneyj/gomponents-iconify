package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamesFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10c-7.732 0-14 6.268-14 14s6.268 14 14 14h12c7.732 0 14-6.268 14-14s-6.268-14-14-14H18Zm-1.75 8c.69 0 1.25.56 1.25 1.25V23h3.25a1.25 1.25 0 1 1 0 2.5H17.5v3.25a1.25 1.25 0 1 1-2.5 0V25.5h-3.75a1.25 1.25 0 1 1 0-2.5H15v-3.75c0-.69.56-1.25 1.25-1.25ZM32 27.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm1.5-4.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}