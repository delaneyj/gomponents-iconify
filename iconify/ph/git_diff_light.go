package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitDiffLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 154a6 6 0 0 0-6 6v33.52l-31.7-31.7a41.75 41.75 0 0 1-12.3-29.7V93.4a30 30 0 1 0-12 0v38.72a53.65 53.65 0 0 0 15.82 38.18l31.7 31.7H64a6 6 0 0 0 0 12h48a6 6 0 0 0 6-6v-48a6 6 0 0 0-6-6ZM38 64a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm168 98.6v-38.72a53.65 53.65 0 0 0-15.82-38.18L158.48 54H192a6 6 0 0 0 0-12h-48a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0V62.48l31.7 31.7a41.75 41.75 0 0 1 12.3 29.7v38.72a30 30 0 1 0 12 0Zm-6 47.4a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}