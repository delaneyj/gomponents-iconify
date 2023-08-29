package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 64v128a8 8 0 0 1-8 8H64a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h128a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M192 48H64a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h128a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16Zm0 144H64V64h128v128Zm48-136v144a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0ZM32 56v144a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}