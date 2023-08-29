package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vncviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.43 7.31L14 38.5a1.53 1.53 0 0 0 3.05 0l9.52-31.19m3.99 13.93v-7.41c0-2.85 3-5.6 6.31-5.6s6.43 3 6.43 5.6v7.41m.13 18.57H36c-2.85 0-5.6-3-5.6-6.31s3-6.43 5.6-6.43h7.41"/>`),
		g.Group(children),
	)
}