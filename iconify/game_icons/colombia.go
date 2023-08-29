package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Colombia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m300.285 18.627l16.062 11.568c-61.797 55.453-66.367 51.932-35.505 133.25l140.122 32.45c-15.192 70.066-3.262 87.445 13.239 129.328l-74.323.29c1.584 53.786 13.364 101.907-6.953 167.86l-24.831-8.94l6.952-46.682l-70.35 6.09c-51.27-68.647-121.674-90.653-186.9-125.28c29.958-14.76 57.869-31.287 67.54-63.568L126.466 143.75c41.01-47.538 75.07-92.922 173.819-125.123z"/>`),
		g.Group(children),
	)
}