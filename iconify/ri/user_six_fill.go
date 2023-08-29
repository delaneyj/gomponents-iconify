package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17c3.663 0 6.866 1.575 8.608 3.925l-1.842.871C17.348 20.116 14.848 19 12.001 19c-2.848 0-5.347 1.116-6.765 2.796l-1.841-.872C5.137 18.574 8.339 17 12 17Zm0-15a5 5 0 0 1 5 5v3a5 5 0 0 1-10 0V7a5 5 0 0 1 5-5Z"/>`),
		g.Group(children),
	)
}