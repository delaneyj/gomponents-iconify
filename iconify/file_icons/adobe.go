package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M322.481 29.346L512 482.654V29.346H322.481zM0 29.346v453.308L189.519 29.346H0zm256.32 166.896l120.797 286.412h-79.18L261.87 391.31H173.3l83.021-195.068z"/>`),
		g.Group(children),
	)
}