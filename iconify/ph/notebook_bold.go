package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 108a12 12 0 0 1 12-12h56a12 12 0 0 1 0 24h-56a12 12 0 0 1-12-12Zm68 28h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Zm52-88v160a20 20 0 0 1-20 20H48a20 20 0 0 1-20-20V48a20 20 0 0 1 20-20h160a20 20 0 0 1 20 20ZM52 204h16V52H52ZM204 52H92v152h112Z"/>`),
		g.Group(children),
	)
}