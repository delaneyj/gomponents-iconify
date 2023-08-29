package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLargeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsTriangleLargeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.767 6.474a.857.857 0 0 0-1.534 0L6.09 40.759A.857.857 0 0 0 6.857 42h34.286a.857.857 0 0 0 .767-1.24L24.767 6.474Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTriangleLargeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}