package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 23.92a9.04 9.04 0 0 0-2.96-6.61l-2.567-2.334A8.365 8.365 0 0 1 3.75 8.799C3.75 4.242 7.444.5 12 .5s8.25 3.741 8.25 8.298c0 2.343-.989 4.6-2.723 6.177l-2.568 2.334a9.041 9.041 0 0 0-2.96 6.61Zm0 0V24v-.057M12 11.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}