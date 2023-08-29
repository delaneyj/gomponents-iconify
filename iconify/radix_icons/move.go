package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.818.932a.45.45 0 0 0-.636 0l-1.75 1.75a.45.45 0 1 0 .636.636L7 2.386V5.5a.5.5 0 0 0 1 0V2.386l.932.932a.45.45 0 0 0 .636-.636L7.818.932ZM8 9.5a.5.5 0 0 0-1 0v3.114l-.932-.932a.45.45 0 0 0-.636.636l1.75 1.75a.45.45 0 0 0 .636 0l1.75-1.75a.45.45 0 0 0-.636-.636L8 12.614V9.5Zm1-2a.5.5 0 0 1 .5-.5h3.114l-.932-.932a.45.45 0 0 1 .636-.636l1.75 1.75a.45.45 0 0 1 0 .636l-1.75 1.75a.45.45 0 0 1-.636-.636L12.614 8H9.5a.5.5 0 0 1-.5-.5ZM3.318 6.068L2.386 7H5.5a.5.5 0 0 1 0 1H2.386l.932.932a.45.45 0 0 1-.636.636l-1.75-1.75a.45.45 0 0 1 0-.636l1.75-1.75a.45.45 0 1 1 .636.636Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}