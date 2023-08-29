package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEntryForPedestrians(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l8.903 8.903M23.5 23.5L13.118 13.118M16 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l.88-4.25l.023-.097m6.097.785C17 11.5 19 12 21 12m-10.5 5.5c-1 3-3 5-5.5 5m.155-9.5l.535-2.583A4.919 4.919 0 0 1 6.22 9M7.5 7.523a4.902 4.902 0 0 1 3-1.023h1.544a2 2 0 0 1 1.958 2.405L13.155 13l-.037.118M9.403 9.403l3.715 3.715M13.35 4.5s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0c0 1.25-1.595 2.25-1.595 2.25h-.3Z"/>`),
		g.Group(children),
	)
}