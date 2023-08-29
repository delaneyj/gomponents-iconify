package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m144.624 65.392l735.744 446.592l-736.736 446.624zm0-64a63.765 63.765 0 0 0-31.088 8.063c-20.32 11.28-32.912 32.705-32.912 55.937l-.992 893.216a63.958 63.958 0 0 0 32.912 55.936a63.937 63.937 0 0 0 31.088 8.065c11.712 0 23.472-3.216 33.775-9.664l736.72-446.624a63.94 63.94 0 0 0 30.257-54.336c0-22.112-11.44-42.672-30.257-54.352L178.4 11.025a64.084 64.084 0 0 0-33.775-9.632z"/>`),
		g.Group(children),
	)
}