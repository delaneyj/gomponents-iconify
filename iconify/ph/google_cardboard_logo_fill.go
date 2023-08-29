package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleCardboardLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h60.69a15.86 15.86 0 0 0 11.31-4.69l21.17-21.17a4 4 0 0 1 5.66 0L152 203.32a15.89 15.89 0 0 0 11.31 4.68H224a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM80 152a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm96 0a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}