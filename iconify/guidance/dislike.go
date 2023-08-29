package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dislike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 15V3m17 .5v.16a16 16 0 0 0 3.761 10.307l.239.283v.25h-9V18a3.5 3.5 0 0 1-3.5 3.5h-.5V19A4.5 4.5 0 0 0 5 14.5h-.5v-11h14Z"/>`),
		g.Group(children),
	)
}