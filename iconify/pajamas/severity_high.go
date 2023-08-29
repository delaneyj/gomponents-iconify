package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeverityHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.713.295l4.992 4.992a1.008 1.008 0 0 1 0 1.426l-4.992 4.992a1.008 1.008 0 0 1-1.426 0L.295 6.713a1.008 1.008 0 0 1 0-1.426L5.287.295a1.008 1.008 0 0 1 1.426 0Z"/>`),
		g.Group(children),
	)
}