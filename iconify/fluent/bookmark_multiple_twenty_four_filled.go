package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 6.748a2.25 2.25 0 0 1 2.25-2.25h9a2.25 2.25 0 0 1 2.25 2.25V21.25a.75.75 0 0 1-1.127.648l-5.623-3.28l-5.622 3.28A.75.75 0 0 1 4 21.25V6.748ZM15.25 2A4.75 4.75 0 0 1 20 6.75v11.873a.75.75 0 0 1-1.5 0V6.751a3.25 3.25 0 0 0-3.25-3.25H6.637s.113-.558.797-1.082C8 2 8.602 2 8.602 2h6.648Z"/>`),
		g.Group(children),
	)
}