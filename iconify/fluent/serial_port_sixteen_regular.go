package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SerialPortSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm1.5 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM8 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM11 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM3.022 5a2 2 0 0 0-1.944 2.47l.725 3A2 2 0 0 0 3.747 12h8.519a2 2 0 0 0 1.944-1.53l.726-3A2 2 0 0 0 12.992 5h-9.97ZM2.05 7.235A1 1 0 0 1 3.022 6h9.97a1 1 0 0 1 .972 1.235l-.726 3a1 1 0 0 1-.972.765h-8.52a1 1 0 0 1-.971-.765l-.725-3Z"/>`),
		g.Group(children),
	)
}