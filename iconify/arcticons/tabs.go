package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tabs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.75 24.38H41.5c0 9.81-7.45 17.75-17.25 17.75a17.75 17.75 0 0 1 0-35.5a17.45 17.45 0 0 1 16.28 11.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.26 8.53c-1.51-1.25-3.45-3.94-3.67-5.4c0 0-3 4.15 0 7.87m19.3-2.6c1.49-1.3 3.31-3.86 3.52-5.27c0 0 2.92 4 .14 7.69"/>`),
		g.Group(children),
	)
}