package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningDiamondDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m229.67 133.62l-96 96a7.94 7.94 0 0 1-11.24 0l-96-96a7.94 7.94 0 0 1 0-11.24l96.05-96a7.94 7.94 0 0 1 11.24 0l96 96.05a7.94 7.94 0 0 1-.05 11.19Z" opacity=".2"/><path d="M128 72a8 8 0 0 1 8 8v56a8 8 0 0 1-16 0V80a8 8 0 0 1 8-8Zm-12 100a12 12 0 1 0 12-12a12 12 0 0 0-12 12Zm124-44a15.85 15.85 0 0 1-4.67 11.28l-96.05 96.06a16 16 0 0 1-22.56 0l-96-96.06a16 16 0 0 1 0-22.56l96.05-96.06a16 16 0 0 1 22.56 0l96.05 96.06A15.85 15.85 0 0 1 240 128Zm-16 0l-96-96l-96 96l96 96Z"/></g>`),
		g.Group(children),
	)
}