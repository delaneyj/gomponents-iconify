package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitDiffBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 148a12 12 0 0 0-12 12v19l-21.46-21.43A35.76 35.76 0 0 1 68 132.12V97.94a36 36 0 1 0-24 0v34.18a59.61 59.61 0 0 0 17.57 42.42L83 196H64a12 12 0 0 0 0 24h48a12 12 0 0 0 12-12v-48a12 12 0 0 0-12-12ZM56 52a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm156 106.06v-34.18a59.61 59.61 0 0 0-17.57-42.42L173 60h19a12 12 0 0 0 0-24h-48a12 12 0 0 0-12 12v48a12 12 0 0 0 24 0V77l21.46 21.46A35.76 35.76 0 0 1 188 123.88v34.18a36 36 0 1 0 24 0ZM200 204a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}