package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.203 1.043a.5.5 0 0 0-.635.709L3.921 7.5L.568 13.248a.5.5 0 0 0 .635.709l13.5-6a.5.5 0 0 0 0-.914l-13.5-6ZM4.846 7.1L2.212 2.586L13.27 7.5L2.212 12.414L4.846 7.9H9a.4.4 0 1 0 0-.8H4.846Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}