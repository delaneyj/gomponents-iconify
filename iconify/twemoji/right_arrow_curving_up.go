package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="m20.589 5.8l6 7.2H23v7.2c0 5.523-4.683 10-10.206 10c-1.414 0-2.861-.298-4.081-.827l2.699-3.3c.396.083.987.127 1.408.127c3.314 0 6.18-2.686 6.18-6V13h-4.411l6-7.2z"/>`),
		g.Group(children),
	)
}