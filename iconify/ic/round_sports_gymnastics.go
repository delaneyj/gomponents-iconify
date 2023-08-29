package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSportsGymnastics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2zm9 16c-.56 0-1.02-.44-1.05-1l-.45-9L8 11H2c-.55 0-1-.45-1-1s.45-1 1-1h5l6.26-4.47c.42-.3 1-.23 1.34.16c.38.45.3 1.12-.18 1.46L11.14 8.5H14l7.09-4.09a.98.98 0 0 1 1.1 1.62L14.5 12l-.45 9c-.03.56-.49 1-1.05 1z"/>`),
		g.Group(children),
	)
}