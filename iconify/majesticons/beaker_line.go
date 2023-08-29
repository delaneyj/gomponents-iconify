package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 4a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2v5.62l4.788 5.387c1.72 1.935.347 4.993-2.242 4.993H5.454c-2.589 0-3.962-3.058-2.243-4.993L8 10.62V5a1 1 0 0 1-1-1zm3 1v6a1 1 0 0 1-.253.664l-5.04 5.672C4.132 17.98 4.59 19 5.453 19h13.092c.863 0 1.321-1.02.748-1.664l-5.041-5.672A1 1 0 0 1 14 11V5h-4z"/></g>`),
		g.Group(children),
	)
}