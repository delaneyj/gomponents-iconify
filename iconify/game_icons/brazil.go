package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brazil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M292.8 41.71c16.1 58.89 125.3 78.19 197.9 116.19c1.6 35.2-14.4 72.6-56.7 102.3c2.9 70.2-41.8 110.2-114.3 132.4c-.3 33.2-12.7 64-47.3 90.3l-59-36.4l47.4-34.2c-1.8-25.6-9.6-52.3-55-67.3l-26.3-93.2c-54.5-10.4-51.9-31.3-56.3-50.9l-64.93 20.4c-49.154-31-51.902-75.4 6.26-83.4l6.99-72.78l51.18 9.12L133 37.03l49.6-7.9l20.7 37.33z"/>`),
		g.Group(children),
	)
}