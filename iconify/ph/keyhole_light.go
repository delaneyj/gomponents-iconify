package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyholeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm38-106a38 38 0 1 0-63.93 27.78L90.71 168.3A10 10 0 0 0 100 182h56a10 10 0 0 0 9.29-13.7l-11.36-28.52A38.14 38.14 0 0 0 166 112Zm-22.78 21.08a6 6 0 0 0-2.06 7.08L153.05 170H103l11.89-29.84a6 6 0 0 0-2.06-7.08a26 26 0 1 1 30.44 0Z"/>`),
		g.Group(children),
	)
}