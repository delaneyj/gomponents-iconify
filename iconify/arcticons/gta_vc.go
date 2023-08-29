package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GtaVc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.135 13.192l5.08 20.488a.677.677 0 0 0 1.25.161L27.206 7m10.812 6.02c4.449.456 3.707-10.442-5.69-3.536c-9.13 6.71-14.101 22.654-2 22.654c7.52 0 11.398-6.684 11.398-6.684M4.418 39.11c10.001-2.736 35.693-4.652 37.665-4.318s1.45 2.456.41 3.546"/>`),
		g.Group(children),
	)
}