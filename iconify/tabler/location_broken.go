package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.896 19.792L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.487 9.657M21.5 21.5l-5-5m0 5l5-5"/>`),
		g.Group(children),
	)
}