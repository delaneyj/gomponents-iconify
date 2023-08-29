package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorArrowRightTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a3 3 0 0 0-3 3v18a3 3 0 0 0 3 3h7.4A7.5 7.5 0 0 1 23 13.427V5a3 3 0 0 0-3-3H8Zm1.5 13a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM27 20.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16.5 20a.5.5 0 0 0 0 1h6.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 .146-.351v-.006a.499.499 0 0 0-.146-.35l-3-3a.5.5 0 0 0-.708.707L23.293 20H16.5Z"/>`),
		g.Group(children),
	)
}