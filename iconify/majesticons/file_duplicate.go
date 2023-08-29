package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDuplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16.83 20H8a3 3 0 0 1-3-3V6.17A3.001 3.001 0 0 0 3 9v10a3 3 0 0 0 3 3h8a3.001 3.001 0 0 0 2.83-2zM7 5v12a1 1 0 0 0 1 1h10a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H10a3 3 0 0 0-3 3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}