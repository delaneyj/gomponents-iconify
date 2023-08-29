package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2v3.866l5.336-3.24l1.038 1.71L14 8.206v8.588l6.374 3.87l-1.038 1.71L14 19.134V22h-2v-4.08L3.073 12.5L12 7.08V2h2Zm-2 7.42L6.927 12.5L12 15.58V9.42Z"/>`),
		g.Group(children),
	)
}