package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.95 1.5a.45.45 0 0 0-.9 0v12a.45.45 0 1 0 .9 0v-12ZM6.568 3.932a.45.45 0 1 0-.636.636L8.414 7.05H.5a.45.45 0 0 0 0 .9h7.914l-2.482 2.482a.45.45 0 1 0 .636.636l3.25-3.25a.45.45 0 0 0 0-.636l-3.25-3.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}