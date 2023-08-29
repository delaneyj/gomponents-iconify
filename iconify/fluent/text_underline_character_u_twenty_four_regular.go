package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineCharacterUTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 4.75a.75.75 0 0 0-1.5 0V12a5 5 0 0 0 10 0V4.75a.75.75 0 0 0-1.5 0V12a3.5 3.5 0 1 1-7 0V4.75ZM6.75 18.5a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H6.75Z"/>`),
		g.Group(children),
	)
}