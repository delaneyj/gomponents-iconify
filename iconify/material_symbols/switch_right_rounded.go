package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 15.375L18.875 12L15.5 8.625v6.75Zm.2 1.925q-.475.475-1.088.213T14 16.575v-9.15q0-.675.613-.937T15.7 6.7l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.6 4.6Zm-7.4 0l-4.6-4.6q-.15-.15-.213-.325T3.426 12q0-.2.063-.375T3.7 11.3l4.6-4.6q.475-.475 1.088-.212t.612.937v9.15q0 .675-.613.938T8.3 17.3Z"/>`),
		g.Group(children),
	)
}