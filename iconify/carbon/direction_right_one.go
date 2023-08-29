package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19 4l-1.414 1.414L22.172 10H10a2 2 0 0 0-2 2v16h2V12h12.172l-4.586 4.586L19 18l7-7Z"/>`),
		g.Group(children),
	)
}