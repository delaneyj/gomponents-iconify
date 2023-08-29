package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m80 96l-40 40V56Z" opacity=".2"/><path d="M224 128a8 8 0 0 1-8 8H112a8 8 0 0 1 0-16h104a8 8 0 0 1 8 8ZM112 72h104a8 8 0 0 0 0-16H112a8 8 0 0 0 0 16Zm104 112H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16ZM32 136V56a8 8 0 0 1 13.66-5.66l40 40a8 8 0 0 1 0 11.32l-40 40A8 8 0 0 1 32 136Zm16-19.31L68.69 96L48 75.31Z"/></g>`),
		g.Group(children),
	)
}