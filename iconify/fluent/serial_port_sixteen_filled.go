package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SerialPortSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.078 7.47A2 2 0 0 1 3.022 5h9.97a2 2 0 0 1 1.944 2.47l-.726 3A2 2 0 0 1 12.265 12h-8.52a2 2 0 0 1-1.943-1.53l-.725-3ZM5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm1.5 1.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM8 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM7 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM11 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}