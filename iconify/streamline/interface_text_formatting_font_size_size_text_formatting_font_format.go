package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingFontSizeSizeTextFormattingFontFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 12.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33M7.91 9.21h4.18M.5 12.5L5.07 1.84a.57.57 0 0 1 1 0L7 3.94M2.55 7.72H5"/>`),
		g.Group(children),
	)
}