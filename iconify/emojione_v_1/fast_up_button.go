package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastUpButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M64 57.1a6.9 6.9 0 0 1-6.9 6.904H6.9A6.9 6.9 0 0 1 .004 57.1V6.9C.004 3.09 3.092 0 6.9 0h50.2A6.9 6.9 0 0 1 64 6.9v50.2"/><g fill="#fff"><path d="M48.48 52.25c-.156 2.317-1.171 4.253-2.557 5.101H18.249c-1.39-.852-2.405-2.791-2.559-5.112l16.337-22.434L48.48 52.25"/><path d="M48.48 29.27c-.156 2.318-1.171 4.253-2.557 5.101H18.249c-1.39-.852-2.405-2.791-2.559-5.112L32.028 6.824L48.48 29.27"/></g>`),
		g.Group(children),
	)
}