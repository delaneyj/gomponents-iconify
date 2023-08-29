package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkpure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.1 4.5L4.73 33.8l4.76 9.7l19.12-39H19.1zm24.17 28.7l-5.05 10.3l-5.05-10.3h10.1zM28.61 4.5l9.32 19l-4.76 9.7l-9.31-19l4.75-9.7zM19.05 24h9.61m-14.12 9.21h18.63"/>`),
		g.Group(children),
	)
}