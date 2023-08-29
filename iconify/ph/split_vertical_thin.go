package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 152a4 4 0 0 1-4 4h-76v66.34l25.17-25.17a4 4 0 0 1 5.66 5.66l-32 32a4 4 0 0 1-5.66 0l-32-32a4 4 0 0 1 5.66-5.66L124 222.34V156H48a4 4 0 0 1 0-8h160a4 4 0 0 1 4 4ZM48 108h160a4 4 0 0 0 0-8h-76V33.66l25.17 25.17a4 4 0 1 0 5.66-5.66l-32-32a4 4 0 0 0-5.66 0l-32 32a4 4 0 0 0 5.66 5.66L124 33.66V100H48a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}