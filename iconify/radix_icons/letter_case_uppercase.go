package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCaseUppercase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.626 2.75a.5.5 0 0 1 .468.327l3.076 8.32a.5.5 0 0 1-.938.346L5.224 9.016H2.027L1.02 11.743a.5.5 0 1 1-.938-.347l3.076-8.32a.5.5 0 0 1 .469-.326Zm0 1.942L4.91 8.166H2.34l1.284-3.474Zm7.746-1.942a.5.5 0 0 1 .469.327l3.075 8.32a.5.5 0 1 1-.938.346L12.97 9.016H9.773l-1.008 2.727a.5.5 0 1 1-.938-.347l3.076-8.32a.5.5 0 0 1 .469-.326Zm0 1.942l1.284 3.474h-2.568l1.284-3.474Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}