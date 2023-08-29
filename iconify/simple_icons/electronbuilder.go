package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electronbuilder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.01a3.506 3.506 0 0 0 3.506-3.505A3.506 3.506 0 0 0 12 0a3.506 3.506 0 0 0-3.506 3.506A3.506 3.506 0 0 0 12 7.01m0 4.137C9.243 8.588 5.574 7.01 1.484 7.01v12.852C5.574 19.863 9.243 21.44 12 24c2.757-2.56 6.426-4.137 10.516-4.137V7.01c-4.09 0-7.759 1.578-10.516 4.137z"/>`),
		g.Group(children),
	)
}