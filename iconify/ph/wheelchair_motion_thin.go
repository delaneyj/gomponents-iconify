package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairMotionThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 76a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm0-48a20 20 0 1 1-20 20a20 20 0 0 1 20-20Zm-12 140a60 60 0 1 1-60-60a4 4 0 0 1 0 8a52 52 0 1 0 52 52a4 4 0 0 1 8 0Zm39.09-34.54a4 4 0 0 1 .83 3.32l-16 80A4 4 0 0 1 184 220a3.44 3.44 0 0 1-.78-.08a4 4 0 0 1-3.14-4.7l15-75.22H128a4 4 0 0 1-3.47-6l22.08-38.42a84.05 84.05 0 0 0-96.06 7.61a4 4 0 0 1-5.1-6.19a92 92 0 0 1 108.73-6.15a4 4 0 0 1 1.29 5.34L134.91 132H200a4 4 0 0 1 3.09 1.46Z"/>`),
		g.Group(children),
	)
}