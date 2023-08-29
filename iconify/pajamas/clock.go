package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0ZM8.75 3.75a.75.75 0 0 0-1.5 0v4.56l.22.22l2.254 2.254a.75.75 0 1 0 1.06-1.06L8.75 7.689V3.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}