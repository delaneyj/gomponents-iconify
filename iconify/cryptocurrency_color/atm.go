package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#346fce"/><path fill="#fff" d="m20.16 18.706l.956 1.797a4.369 4.369 0 0 1-1.806 5.908l-.124.066a4.369 4.369 0 0 1-5.909-1.806l-3.223-6.062a4.369 4.369 0 0 1 1.592-5.787l-.956-1.797a4.369 4.369 0 0 1 1.807-5.908l.124-.066a4.369 4.369 0 0 1 5.908 1.806l3.223 6.062a4.369 4.369 0 0 1-1.591 5.787zm0 0l-2.267-4.265a4.369 4.369 0 0 0-5.908-1.806l-.124.065c-.073.04-.145.08-.215.122l2.268 4.265a4.369 4.369 0 0 0 5.908 1.806l.124-.065c.073-.04.145-.08.215-.122z"/></g>`),
		g.Group(children),
	)
}