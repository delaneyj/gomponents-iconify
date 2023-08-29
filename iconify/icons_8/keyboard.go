package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 7c-1.093 0-2 .907-2 2v14c0 1.093.907 2 2 2h26c1.093 0 2-.907 2-2V9c0-1.093-.907-2-2-2H3zm0 2h26v14H3V9zm2 2v2h2v-2H5zm4 0v2h2v-2H9zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zM5 15v2h4v-2H5zm6 0v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h4v-2h-4zM5 19v2h4v-2H5zm6 0v2h10v-2H11zm12 0v2h4v-2h-4z"/>`),
		g.Group(children),
	)
}