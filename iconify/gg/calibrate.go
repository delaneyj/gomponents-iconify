package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calibrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.05 5a8.97 8.97 0 0 1 6.314 2.586l-4.243 4.243A2.99 2.99 0 0 0 12.05 11c-.855 0-1.625.357-2.172.93L5.636 7.687A8.973 8.973 0 0 1 12.05 5Zm0 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}