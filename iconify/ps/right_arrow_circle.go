package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469zm107-219q-3 0 0 0q0-3-2-5q0-2-3-2l-42-64q-12-17-30-6q-18 10-7 30l22 30H171q-22 0-22 21t22 21h130l-22 30q-11 20 7 30q19 10 30-6l42-64q0-3 3-3v-10q2 2 2-2z"/>`),
		g.Group(children),
	)
}