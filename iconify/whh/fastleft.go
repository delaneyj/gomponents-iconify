package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M435 467q-19 18-19 44.5t19 45.5l451 448q19 19 46 19t46-19l27-27q19-18 19-44.5t-19-45.5L626 512l379-376q19-19 19-45.5T1005 45l-27-26Q959 0 932 0t-46 19zm154 511q19-18 19-44.5T589 888L210 512l379-376q19-19 19-45.5T589 45l-27-26Q543 0 516.5 0T471 19L19 467Q0 485 0 511.5T19 557l452 448q19 19 45.5 19t45.5-19z"/>`),
		g.Group(children),
	)
}