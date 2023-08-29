package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.75 1.5a.75.75 0 0 0-1.5 0v1.03a9.5 9.5 0 0 0-8.749 9.319a.667.667 0 0 0 .558.668a54.39 54.39 0 0 0 8.191.735V19a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 0-1.5 0v.5a2.25 2.25 0 0 0 4.5 0v-5.748a54.39 54.39 0 0 0 8.192-.735a.667.667 0 0 0 .557-.668c-.077-4.926-3.902-8.941-8.749-9.32V1.5Z"/>`),
		g.Group(children),
	)
}