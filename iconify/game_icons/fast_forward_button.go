package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M44.62 106a26 26 0 0 0-25.69 29.3c16 124 16 117.4 0 241.4A26 26 0 0 0 54.72 404l138.68-57.7c-1.2 9.5-2.4 18.9-3.9 30.4c-2.5 19.8 17.3 35 35.8 27.3l252-124c9.7-4 16-13.5 16-24s-6.3-20-16-24l-252-124c-3.2-1.3-6.6-2-10.1-2c-15.6.1-27.7 13.8-25.7 29.3c1.5 11.5 2.7 20.9 3.9 30.4L54.72 108a26 26 0 0 0-10.1-2z"/>`),
		g.Group(children),
	)
}