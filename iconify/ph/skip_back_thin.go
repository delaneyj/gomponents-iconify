package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M197.87 37.52a12 12 0 0 0-12.18.32L60 116.44V40a4 4 0 0 0-8 0v176a4 4 0 0 0 8 0v-76.44l125.69 78.6a12 12 0 0 0 12.18.32a11.69 11.69 0 0 0 6.13-10.36V47.88a11.69 11.69 0 0 0-6.13-10.36ZM196 208.12a3.8 3.8 0 0 1-2 3.37a4 4 0 0 1-4.06-.11L61.82 131.26a3.83 3.83 0 0 1 0-6.52l128.12-80.12a4 4 0 0 1 2.11-.62a4 4 0 0 1 1.95.51a3.8 3.8 0 0 1 2 3.37Z"/>`),
		g.Group(children),
	)
}