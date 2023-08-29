package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h12v5"/><path d="M13 16H6a2 2 0 0 0-2 2m11 1l3 3l3-3m-3 3v-9"/></g>`),
		g.Group(children),
	)
}