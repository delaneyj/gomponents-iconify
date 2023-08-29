package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 4L8 6l2 2l2-2zm2 2l2 2l2-2l-2-2zm4 0l2 2l2-2l-2-2zm4 0l2 2l2-2l-2-2zm2 2H5v20h22V8zM7 10h18v16h-6v-5h-6v5H7zm2 2v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zM9 15v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zM9 18v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zM9 21v2h2v-2zm12 0v2h2v-2zm-6 2h2v3h-2z"/>`),
		g.Group(children),
	)
}