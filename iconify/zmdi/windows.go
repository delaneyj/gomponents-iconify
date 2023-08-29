package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 192V80l128-28v138zM363 0v187l-214 3V47zM0 213l128 2v146L0 336V213zm363 6v186l-214-40V215z"/>`),
		g.Group(children),
	)
}