package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Token(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4.5h6a3.5 3.5 0 1 1 0 7H5a3.5 3.5 0 1 1 0-7ZM0 8a5 5 0 0 1 5-5h6a5 5 0 0 1 0 10H5a5 5 0 0 1-5-5Zm11 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}