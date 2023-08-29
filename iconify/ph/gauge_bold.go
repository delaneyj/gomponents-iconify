package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaugeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M209.88 77.83A115.19 115.19 0 0 0 128 44h-.41C63.85 44.22 12 96.76 12 161.13V184a20 20 0 0 0 20 20h192a20 20 0 0 0 20-20v-24a115.25 115.25 0 0 0-34.12-82.17ZM220 180h-92.68l46.44-65a12 12 0 1 0-19.52-14l-56.42 79H36v-18.87c0-1.72 0-3.43.14-5.13H56a12 12 0 0 0 0-24H40.62c10.91-33.39 40-58.52 75.38-63.21V88a12 12 0 0 0 24 0V68.8a92 92 0 0 1 75.66 63.2H200a12 12 0 0 0 0 24h19.9c.06 1.33.1 2.66.1 4Z"/>`),
		g.Group(children),
	)
}