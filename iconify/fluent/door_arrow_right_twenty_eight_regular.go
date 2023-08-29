package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorArrowRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8.427a7.451 7.451 0 0 0-1.5-.36V5A1.5 1.5 0 0 0 20 3.5H8A1.5 1.5 0 0 0 6.5 5v18A1.5 1.5 0 0 0 8 24.5h6.155a7.543 7.543 0 0 0 1.246 1.5H8a3 3 0 0 1-3-3V5Zm4.5 10a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM27 20.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16.5 20a.5.5 0 0 0 0 1h6.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 .146-.351v-.006a.499.499 0 0 0-.146-.35l-3-3a.5.5 0 0 0-.708.707L23.293 20H16.5Z"/>`),
		g.Group(children),
	)
}