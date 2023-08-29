package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 1 0 0 18.5a9.25 9.25 0 0 0 0-18.5Zm5.48 18.5A10.745 10.745 0 0 0 22.75 12c0-5.937-4.813-10.75-10.75-10.75S1.25 6.063 1.25 12S6.063 22.75 12 22.75h10a.75.75 0 0 0 0-1.5h-4.52ZM12 6.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm-2.25.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM7.5 11.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM5.25 12a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Zm11.25-.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm-2.25.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM12 15.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm-2.25.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}