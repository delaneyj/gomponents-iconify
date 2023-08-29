package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 2v3.252a.75.75 0 0 0 1.5 0V2h1v4.251a.75.75 0 0 0 1.5 0V2h1.75a.75.75 0 0 1 .75.75V11H5V2.75A.75.75 0 0 1 5.75 2h6.75ZM5 12.5v1.752a2.25 2.25 0 0 0 2.25 2.25H10V20a2 2 0 1 0 4 0v-3.498h2.75a2.25 2.25 0 0 0 2.25-2.25V12.5H5Z"/>`),
		g.Group(children),
	)
}