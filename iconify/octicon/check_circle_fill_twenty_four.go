package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleFillTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm16.28-2.72a.751.751 0 0 0-.018-1.042a.751.751 0 0 0-1.042-.018l-5.97 5.97l-2.47-2.47a.751.751 0 0 0-1.042.018a.751.751 0 0 0-.018 1.042l3 3a.75.75 0 0 0 1.06 0Z"/>`),
		g.Group(children),
	)
}