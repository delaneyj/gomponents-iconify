package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M80 40v176H48a8 8 0 0 1-8-8V48a8 8 0 0 1 8-8Z" opacity=".2"/><path d="M184 112a8 8 0 0 1-8 8h-64a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm-8 24h-64a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm48-88v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16ZM48 208h24V48H48Zm160 0V48H88v160h120Z"/></g>`),
		g.Group(children),
	)
}