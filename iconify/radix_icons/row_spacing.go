package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.818.682a.45.45 0 0 0-.636 0l-2 2a.45.45 0 0 0 .636.636L7.05 2.086V5.5a.45.45 0 1 0 .9 0V2.086l1.232 1.232a.45.45 0 0 0 .636-.636l-2-2Zm.132 12.232V9.5a.45.45 0 1 0-.9 0v3.414l-1.232-1.232a.45.45 0 0 0-.636.636l2 2a.45.45 0 0 0 .636 0l2-2a.45.45 0 1 0-.636-.636L7.95 12.914ZM1.5 7a.5.5 0 0 0 0 1h12a.5.5 0 1 0 0-1h-12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}