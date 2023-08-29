package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sushi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C7.51 2 4 4.2 4 7v10c0 2.8 3.51 5 8 5s8-2.2 8-5V7c0-2.8-3.51-5-8-5zm0 18c-3.54 0-6-1.58-6-3v-6.67A10.52 10.52 0 0 0 12 12a10.52 10.52 0 0 0 6-1.67V17c0 1.42-2.46 3-6 3zm0-10c-3.54 0-6-1.58-6-3s2.46-3 6-3s6 1.58 6 3s-2.46 3-6 3z"/><ellipse cx="12" cy="7" fill="currentColor" rx="3" ry="1.71"/>`),
		g.Group(children),
	)
}