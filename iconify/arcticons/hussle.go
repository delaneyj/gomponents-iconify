package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hussle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.177 43.5l-5.232-3.547l10.976-23.831l-2.625 1.751l-3.556-5.329L23.795 4.5l4.687 4.004l-7.624 16.555l7.122-7.497l5.526 2.263l-.225 12.668l3.394-3.177l4.38 4.676l-8.921 8.353l-5.394-2.395l.214-12.005L12.177 43.5z"/>`),
		g.Group(children),
	)
}