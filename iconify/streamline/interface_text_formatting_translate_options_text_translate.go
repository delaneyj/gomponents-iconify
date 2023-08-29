package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTranslateOptionsTextTranslate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 13.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33m-5.59-3.29h4.18M.5 2.5h9M5 .5v2m3 0l-6 7m0-7l3.94 4.6"/>`),
		g.Group(children),
	)
}