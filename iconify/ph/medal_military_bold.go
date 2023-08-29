package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalMilitaryBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M207 28H49a21 21 0 0 0-21 21v49.21a21 21 0 0 0 12.31 19.11l56 25.47a52 52 0 1 0 63.32 0l56-25.47A21 21 0 0 0 228 98.21V49a21 21 0 0 0-21-21Zm-79 102.82l-28-12.73V52h56v66.09ZM52 52h24v55.18L52 96.27Zm76 160a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm76-115.73l-24 10.91V52h24Z"/>`),
		g.Group(children),
	)
}