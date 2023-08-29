package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 208H48l80-80Z" opacity=".2"/><path d="M133.66 122.34a8 8 0 0 0-11.32 0l-80 80A8 8 0 0 0 48 216h160a8 8 0 0 0 5.66-13.66ZM67.31 200L128 139.31L188.69 200Zm-25-66.34a8 8 0 0 1 0-11.32l80-80a8 8 0 0 1 11.32 0l80 80a8 8 0 0 1-11.32 11.32L128 59.31l-74.34 74.35a8 8 0 0 1-11.32 0Z"/></g>`),
		g.Group(children),
	)
}