package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M12 3a4.5 4.5 0 0 0 0 9a4.5 4.5 0 0 1 0 9"/><circle cx="12" cy="7.5" r=".5" fill="currentColor"/><circle cx="12" cy="16.5" r=".5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}