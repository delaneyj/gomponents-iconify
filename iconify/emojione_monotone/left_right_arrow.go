package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M10.787 10.787c-11.715 11.715-11.715 30.711 0 42.426c11.714 11.717 30.71 11.717 42.426 0c11.717-11.715 11.717-30.71 0-42.426c-11.716-11.716-30.712-11.716-42.426 0zM36.487 16l15.512 16l-15.512 16V36.628h-8.975V48L12 32l15.513-16v11.373h8.975V16z"/>`),
		g.Group(children),
	)
}