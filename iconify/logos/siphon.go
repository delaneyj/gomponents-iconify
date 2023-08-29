package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Siphon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M0 256h129V128H0v128z" fill="#4C728C"/><path d="M128 129h128V0H128v129z" fill="#6296BA"/><path d="M128 256h128V128H128v128z" fill="#5884A3"/>`),
		g.Group(children),
	)
}