package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Namespace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.3 4.125a.75.75 0 0 0-1.3-.75l-4.9 8.5a.75.75 0 1 0 1.3.75l4.9-8.5ZM8 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2.5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}