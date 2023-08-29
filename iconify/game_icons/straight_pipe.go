package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightPipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M445.8 26.44c-.6 0-1.1 0-1.6.1c-2 .26-3.6 1.14-5.1 2.63L54.19 414c10.18 1.9 20.27 7.3 28.39 15.4c8.14 8.2 13.47 18.3 15.43 28.5L482.9 73.03c3-3 3.5-6.42 1.7-12.57c-1.9-6.15-6.8-13.73-13-19.95c-6.2-6.23-13.8-11.1-20-12.98c-2.3-.68-4.2-1.06-5.8-1.09zM45.19 431.2c-5.39.2-10.19 2-13.52 5.3c-8.2 8.2-7.21 25.3 5.66 38.2c12.87 12.8 29.98 13.8 38.18 5.6c8.2-8.2 7.21-25.3-5.66-38.2c-7.64-7.6-16.78-11-24.66-10.9z"/>`),
		g.Group(children),
	)
}