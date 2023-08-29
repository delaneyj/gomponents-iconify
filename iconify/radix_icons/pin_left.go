package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.05 13.5a.45.45 0 0 0 .9 0v-12a.45.45 0 1 0-.9 0v12Zm6.382-2.432a.45.45 0 1 0 .636-.636L6.586 7.95H14.5a.45.45 0 0 0 0-.9H6.586l2.482-2.482a.45.45 0 1 0-.636-.636l-3.25 3.25a.45.45 0 0 0 0 .636l3.25 3.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}