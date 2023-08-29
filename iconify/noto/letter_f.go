package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40C0E7" d="M101.15 15.54H34.86c-1.23 0-2.24 1-2.24 2.24v100.97c0 1.24 1.01 2.24 2.24 2.24h21.89c1.24 0 2.24-1 2.24-2.24v-44H94.1c1.24 0 2.24-1 2.24-2.24V55.79c0-1.24-1-2.24-2.24-2.24H58.99v-16.8h42.16c1.24 0 2.24-1 2.24-2.24V17.78c0-1.24-1-2.24-2.24-2.24z"/>`),
		g.Group(children),
	)
}