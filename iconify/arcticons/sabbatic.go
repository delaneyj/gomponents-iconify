package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sabbatic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 20.266l6.953 14.078h25.094L43.5 20.266l-19.499 5.467l-19.5-5.467Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.433 23.05l9.188-9.395l9.366 9.558m3.217 10.813l-2.221-8.648M31.7 34.344l-2.008-7.864m4.291 7.864l-2.146-8.415m6.035 5.688l-1.744-6.79m3.256 3.773l-1.11-4.324M22.59 25.337l4.357-4.262m-9.28 2.89l6.18-6.054m-3.695 6.75l5.245-5.168m.502 5.71l2.598-2.545"/>`),
		g.Group(children),
	)
}