package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFilterTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 2a7.503 7.503 0 0 1 7.179 5.321a7.5 7.5 0 1 1-9.357 9.358A7.5 7.5 0 0 1 9.5 2Zm7.486 7.038l.01.22L17 9.5a7.5 7.5 0 0 1-7.962 7.486a6 6 0 1 0 7.947-7.948ZM9.5 3.5a6 6 0 0 0-2.486 11.463l-.01-.22L7 14.5a7.5 7.5 0 0 1 7.962-7.486A5.999 5.999 0 0 0 9.5 3.5Z"/>`),
		g.Group(children),
	)
}