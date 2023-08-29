package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanjidraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.18 17.41v-5.3a1.9 1.9 0 0 1 1.9-1.9h31.84a1.9 1.9 0 0 1 1.9 1.9v5.3M24 4.5v5.71m-12.57 7.2h24L24 25.63v13.8a4.07 4.07 0 0 1-4.66 4l-5.51-.81M6.18 28.95h35.64"/>`),
		g.Group(children),
	)
}