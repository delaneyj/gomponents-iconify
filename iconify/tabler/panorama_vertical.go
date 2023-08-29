package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.463 4.338c-1.932 5.106-1.932 10.211 0 15.317A1 1 0 0 1 17.529 21h-11c-.692 0-1.208-.692-.962-1.34c1.932-5.107 1.932-10.214 0-15.321C5.321 3.691 5.81 3 6.502 3H17.53c.693 0 1.18.691.935 1.338z"/>`),
		g.Group(children),
	)
}