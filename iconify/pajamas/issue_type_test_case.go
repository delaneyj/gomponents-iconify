package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeTestCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.751 7.834L7.5 7.4V3.5h1v3.9l.749.434A2.498 2.498 0 0 1 10.5 10h-5c0-.923.5-1.732 1.251-2.166z"/><path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4zm6 3.5v3.035a4 4 0 1 1-4 0V3.5h-.25a.75.75 0 0 1 0-1.5h4.5a.75.75 0 0 1 0 1.5H10z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}