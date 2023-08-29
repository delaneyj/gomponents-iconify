package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-80 128a8 8 0 0 1-16 0v-32H72v32a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0v32h40V80a8 8 0 0 1 16 0Zm40 24a32 32 0 0 1-21.34-8.15a8 8 0 1 1 10.68-11.92A16 16 0 1 0 168 136a8 8 0 0 1-6.4-12.8L176 104h-24a8 8 0 0 1 0-16h40a8 8 0 0 1 6.4 12.8l-16.71 22.28A32 32 0 0 1 168 184Z"/>`),
		g.Group(children),
	)
}