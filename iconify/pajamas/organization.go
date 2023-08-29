package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organization(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0Zm-8 2.25a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM4.516 8a3 3 0 1 0 4.932 2.81a3 3 0 1 0 0-5.62A3 3 0 1 0 4.516 8ZM6.5 7.25a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm4 2.25a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}