package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAppleArcade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10 7.5v4.75a.734.734 0 0 1-.055.325a.704.704 0 0 1-.348.366l-5.462 2.58a5 5 0 0 1-4.27 0l-5.462-2.58a.705.705 0 0 1-.401-.691V12.5"/><path d="m4.431 12.216l5.634-2.332a5.065 5.065 0 0 1 3.87 0l5.634 2.332a.692.692 0 0 1 .028 1.269l-5.462 2.543a5.064 5.064 0 0 1-4.27 0l-5.462-2.543a.691.691 0 0 1 .028-1.27zM12 7v6"/></g>`),
		g.Group(children),
	)
}