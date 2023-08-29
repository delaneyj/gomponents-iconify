package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v4h2V7h2V5H5zm6 0v2h4V5h-4zm6 0v2h4V5h-4zm6 0v2h2v2h2V5h-4zm-8 4v4h2V9h-2zM5 11v4h2v-4H5zm20 0v4h2v-4h-2zM9 15v2h4v-2H9zm6 0v2h2v-2h-2zm4 0v2h4v-2h-4zM5 17v4h2v-4H5zm20 0v4h2v-4h-2zm-10 2v4h2v-4h-2zM5 23v4h4v-2H7v-2H5zm20 0v2h-2v2h4v-4h-2zm-14 2v2h4v-2h-4zm6 0v2h4v-2h-4z"/>`),
		g.Group(children),
	)
}