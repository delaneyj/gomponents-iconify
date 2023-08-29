package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CabinCareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 25v-2h-9.232l-.8-3H27v-2h-7.565l-2.47-9.258l-1.93.516l3.8 14.257A2.004 2.004 0 0 0 20.77 25H22v3H10v-3h6v-2H6.768l-.8-3H13v-2H5.435l-2.47-9.258l-1.93.516l3.8 14.257A2.004 2.004 0 0 0 6.77 25H8v3H2v2h28v-2h-6v-3zM27.303 2a2.661 2.661 0 0 0-1.908.806L25 3.21l-.395-.405a2.661 2.661 0 0 0-3.816 0a2.8 2.8 0 0 0 0 3.896L25 11l4.21-4.298a2.8 2.8 0 0 0 0-3.896A2.661 2.661 0 0 0 27.304 2z"/>`),
		g.Group(children),
	)
}