package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleNineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M145 78.55A34 34 0 1 0 127.94 142a33.56 33.56 0 0 0 5.67-.49l-18.76 31.42a6 6 0 0 0 10.3 6.16l32.3-54.09A34 34 0 0 0 145 78.55Zm2.06 40.45A22 22 0 1 1 139 89a22 22 0 0 1 8.05 30ZM128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Z"/>`),
		g.Group(children),
	)
}