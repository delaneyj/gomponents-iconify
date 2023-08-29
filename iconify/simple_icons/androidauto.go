package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Androidauto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0c-.6 0-1.11.32-1.39.8L.48 18.4a1.6 1.6 0 0 0 1.39 2.4h2l7.7-13.58l.43-.77l8.13 14.35h2a1.6 1.6 0 0 0 1.39-2.4L13.39.8A1.6 1.6 0 0 0 12 0zm0 7.47l-9.07 16l.54.53L12 20.8l8.53 3.2l.54-.53z"/>`),
		g.Group(children),
	)
}