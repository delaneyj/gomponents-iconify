package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FF9800" d="M1 38h46v2H1zm24-20h4v16h-4zm6 0h4v16h-4zm6 0h4v16h-4zm-18 0h4v16h-4zm-6 0h4v16h-4zm-6 0h4v16H7zm36-2H5v-3l19-9l19 9zM5 34h38v2H5z"/><g fill="#EF6C00"><path d="M25 16h4v2h-4zm6 0h4v2h-4zm6 0h4v2h-4zm-18 0h4v2h-4zm-6 0h4v2h-4zm-6 0h4v2H7zM3 36h42v2H3z"/><circle cx="24" cy="11" r="2"/></g>`),
		g.Group(children),
	)
}