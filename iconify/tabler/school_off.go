package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 9L12 5l-2.136.854M7 7L2 9l10 4l.697-.279m2.878-1.151L22 9v6"/><path d="M6 10.6V16c0 1.657 2.686 3 6 3c2.334 0 4.357-.666 5.35-1.64M18 14v-3.4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}