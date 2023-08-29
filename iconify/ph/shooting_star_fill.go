package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShootingStarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m235.39 84.07l-28.15 24l8.43 35.73a13.09 13.09 0 0 1-5 13.58a13.25 13.25 0 0 1-14.63.7l-32-19l-32 19a13.25 13.25 0 0 1-14.63-.7a13.1 13.1 0 0 1-5-13.58l8.43-35.73l-28.16-24A13.13 13.13 0 0 1 100.1 61l37.23-3.15L151.85 24a13.24 13.24 0 0 1 24.31 0l14.52 33.87L227.9 61a13.12 13.12 0 0 1 7.49 23.06ZM85.66 114.34a8 8 0 0 0-11.32 0l-56 56a8 8 0 0 0 11.32 11.32l56-56a8 8 0 0 0 0-11.32Zm16 56a8 8 0 0 0-11.32 0l-56 56a8 8 0 0 0 11.32 11.32l56-56a8 8 0 0 0 0-11.32Zm60.69 0l-56 56a8 8 0 0 0 11.32 11.32l56-56a8 8 0 0 0-11.31-11.32Z"/>`),
		g.Group(children),
	)
}