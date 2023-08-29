package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracesCheckmarkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v3a.5.5 0 0 1-.5.5a.5.5 0 0 0 0 1A.5.5 0 0 1 2 9v3a2 2 0 0 0 2 2a.5.5 0 0 0 0-1a1 1 0 0 1-1-1V9c0-.384-.144-.735-.382-1C2.856 7.735 3 7.384 3 7V4a1 1 0 0 1 1-1a.5.5 0 0 0 0-1Zm8 0a2 2 0 0 1 2 2v3a.5.5 0 0 0 .5.5a.5.5 0 0 1 0 1a.5.5 0 0 0-.5.5v3a2 2 0 0 1-2 2a.5.5 0 0 1 0-1a1 1 0 0 0 1-1V9c0-.384.144-.735.382-1A1.494 1.494 0 0 1 13 7V4a1 1 0 0 0-1-1a.5.5 0 0 1 0-1ZM8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm1.146-5.354a.5.5 0 1 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-1-1a.5.5 0 1 1 .708-.708l.646.647l1.646-1.647Z"/>`),
		g.Group(children),
	)
}