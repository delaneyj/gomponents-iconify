package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Algeria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.4 23.98c-87.5-7.58-164.6 2.58-228 35.99l11.3 83.23l-181.56 71.3l-5.99 43.4L294.5 490.4l57-2.2l146.4-114c-60.4-73.4-22.3-118.9-45.6-161.2c-48-86.8-45.2-133.74-33.9-189.02z"/>`),
		g.Group(children),
	)
}