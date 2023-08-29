package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182 112a6 6 0 0 1-6 6h-64a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm-6 26h-64a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm46-90v160a14 14 0 0 1-14 14H48a14 14 0 0 1-14-14V48a14 14 0 0 1 14-14h160a14 14 0 0 1 14 14ZM48 210h26V46H48a2 2 0 0 0-2 2v160a2 2 0 0 0 2 2ZM210 48a2 2 0 0 0-2-2H86v164h122a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}