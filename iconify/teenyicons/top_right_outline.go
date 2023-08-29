package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 1.5h.5a.5.5 0 0 0-.5-.5v.5Zm0-.5H8v1h5.5V1Zm-.5.5V7h1V1.5h-1Zm.146-.354l-12 12l.708.708l12-12l-.708-.708Z"/>`),
		g.Group(children),
	)
}