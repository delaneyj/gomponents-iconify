package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionStraightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-6.414 11.414L17 8.828V26h-2V8.828l-4.586 4.586L9 12l7-7l7 7Z"/><path fill="none" d="M21.586 13.414L17 8.828V26h-2V8.828l-4.586 4.586L9 12l7-7l7 7Z"/>`),
		g.Group(children),
	)
}