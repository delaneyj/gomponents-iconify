package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a4.5 4.5 0 0 0-3.155 7.708a.653.653 0 0 1 .18.287L5.3 11h5.4l.274-1.005a.654.654 0 0 1 .181-.287A4.5 4.5 0 0 0 8 2Zm2.427 10H5.573l.244.895A1.5 1.5 0 0 0 7.264 14h1.472a1.5 1.5 0 0 0 1.447-1.105l.244-.895Z"/>`),
		g.Group(children),
	)
}