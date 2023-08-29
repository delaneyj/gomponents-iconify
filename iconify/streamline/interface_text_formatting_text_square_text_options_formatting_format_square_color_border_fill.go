package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextSquareTextOptionsFormattingFormatSquareColorBorderFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 10.78l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33M4.69 8h4.62"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`),
		g.Group(children),
	)
}