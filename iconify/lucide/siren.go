package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Siren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v6H7v-6Zm-2 8a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2H5v-2Zm16-8h1m-3.5-7.5L18 5M2 12h1m9-10v1M4.929 4.929l.707.707M12 12v6"/>`),
		g.Group(children),
	)
}