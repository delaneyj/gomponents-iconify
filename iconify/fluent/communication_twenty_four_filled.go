package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunicationTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 5a8 8 0 0 0-5.662 13.652a1 1 0 1 1-1.416 1.413A9.972 9.972 0 0 1 2 13C2 7.477 6.477 3 12 3s10 4.477 10 10a9.972 9.972 0 0 1-2.922 7.065a1 1 0 0 1-1.416-1.413A8 8 0 0 0 12 5Zm0 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 6.828 2.828a1 1 0 0 0 1.415 1.415a6 6 0 1 0-8.485 0a1 1 0 1 0 1.414-1.415A3.984 3.984 0 0 1 8 13Z"/>`),
		g.Group(children),
	)
}