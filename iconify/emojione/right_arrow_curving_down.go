package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M40.4 45.1V22.9c0-7.1-5.6-12.9-12.5-12.9c-3.3 0-6.5 1.3-8.9 3.8l5.1 5.2c1-1 2.4-1.6 3.8-1.6c3 0 5.4 2.5 5.4 5.5v22.2h-4.6l8.2 8.9l8.2-8.9h-4.7"/>`),
		g.Group(children),
	)
}