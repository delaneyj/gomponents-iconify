package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 10a2 2 0 1 0 4 0V5a2 2 0 1 0-4 0v5ZM5 8a2 2 0 1 0 0 4h5a2 2 0 1 0 0-4H5Zm10 5a2 2 0 1 0 0 4h5a2 2 0 1 0 0-4h-5Zm-5 9a2 2 0 0 1-2-2v-5a2 2 0 1 1 4 0v5a2 2 0 0 1-2 2ZM8 5a2 2 0 1 1 4 0v2h-2a2 2 0 0 1-2-2ZM3 15a2 2 0 1 0 4 0v-2H5a2 2 0 0 0-2 2Zm14 5a2 2 0 1 1-4 0v-2h2a2 2 0 0 1 2 2Zm5-10a2 2 0 1 0-4 0v2h2a2 2 0 0 0 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}