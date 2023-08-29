package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M20 108a12 12 0 0 1 12-12h76a12 12 0 0 1 0 24H32a12 12 0 0 1-12-12Zm204-12h-76a12 12 0 0 0 0 24h76a12 12 0 0 0 0-24ZM68 136H32a12 12 0 0 0 0 24h36a12 12 0 0 0 0-24Zm156 0h-36a12 12 0 0 0 0 24h36a12 12 0 0 0 0-24ZM96 148a12 12 0 0 0 12 12h40a12 12 0 0 0 0-24h-40a12 12 0 0 0-12 12Zm-44 28H32a12 12 0 0 0 0 24h20a12 12 0 0 0 0-24Zm56 0H92a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Zm56 0h-16a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Zm60 0h-20a12 12 0 0 0 0 24h20a12 12 0 0 0 0-24ZM32 80h192a12 12 0 0 0 0-24H32a12 12 0 0 0 0 24Z"/>`),
		g.Group(children),
	)
}