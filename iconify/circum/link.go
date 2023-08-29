package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.9 8a4.055 4.055 0 0 1 1.352.135a2.511 2.511 0 0 1-.7 4.863a.5.5 0 0 0 0 1a3.508 3.508 0 0 0 2.944-5.2A3.557 3.557 0 0 0 11.434 7H5.59a3.5 3.5 0 0 0-.19 7c.724.041 1.458 0 2.183 0a.5.5 0 0 0 0-1c-1.323 0-2.915.262-3.891-.843A2.522 2.522 0 0 1 5.59 8Z"/><path fill="currentColor" d="M18.41 17a3.5 3.5 0 0 0 .192-6.994c-.724-.041-1.458 0-2.183 0a.5.5 0 0 0 0 1c1.323 0 2.915-.262 3.891.843A2.522 2.522 0 0 1 18.41 16H13.1a4.055 4.055 0 0 1-1.352-.135a2.511 2.511 0 0 1 .7-4.863a.5.5 0 0 0 0-1a3.508 3.508 0 0 0-2.944 5.2A3.557 3.557 0 0 0 12.566 17Z"/>`),
		g.Group(children),
	)
}