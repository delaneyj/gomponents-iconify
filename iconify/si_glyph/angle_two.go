package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.648 15.938a.617.617 0 0 1-.436-1.052L14.946.18a.617.617 0 0 1 .871.872L1.085 15.757a.615.615 0 0 1-.437.181zm5.004.072a.615.615 0 0 1-.441-.185a.623.623 0 0 1 0-.882l9.723-9.723a.62.62 0 0 1 .881 0a.62.62 0 0 1 0 .883l-9.723 9.722a.615.615 0 0 1-.44.185zm4.981-.025a.606.606 0 0 1-.426-1.031l4.752-4.752a.602.602 0 0 1 .854 0a.606.606 0 0 1 0 .853l-4.753 4.752a.6.6 0 0 1-.427.178z"/>`),
		g.Group(children),
	)
}