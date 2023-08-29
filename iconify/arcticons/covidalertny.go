package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covidalertny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.73 8.32l-7.64 3l-2.78 6L19.2 23L8.55 26.26l.69 4.86l-4.74 3l2.2 1.38l19.67-2.19l-.37 3.36l10.41 3l6.83-6.13l.23-2.09l-8.45 2.77Z"/>`),
		g.Group(children),
	)
}