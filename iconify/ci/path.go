package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Path(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.121 15.879a3 3 0 1 0-4.243 4.243a3 3 0 0 0 4.243-4.243Zm0 0L15.88 8.12m0 0a3 3 0 1 0 4.243-4.243A3 3 0 0 0 15.88 8.12Zm0 0l.004-.004"/>`),
		g.Group(children),
	)
}