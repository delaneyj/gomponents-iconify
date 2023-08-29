package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderTransgenderBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 28h-44a12 12 0 0 0 0 24h15l-19 19l-15.51-15.48a12 12 0 0 0-17 17L151 88l-10.42 10.46a76 76 0 1 0 17 17L168 105l15.51 15.52a12 12 0 0 0 17-17L185 88l19-19v15a12 12 0 0 0 24 0V40a12 12 0 0 0-12-12Zm-83.23 168.8A52 52 0 1 1 148 160a52.05 52.05 0 0 1-15.23 36.8Z"/>`),
		g.Group(children),
	)
}