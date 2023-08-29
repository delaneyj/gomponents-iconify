package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Child(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 64a64 64 0 1 1 128 0a64 64 0 1 1-128 0zm48 320v96c0 17.7-14.3 32-32 32s-32-14.3-32-32V287.8L59.1 321c-9.4 15-29.2 19.4-44.1 10s-19.5-29.1-10.1-44l39.9-63.3C69.7 184 113.2 160 160 160s90.3 24 115.2 63.6l39.9 63.4c9.4 15 4.9 34.7-10 44.1s-34.7 4.9-44.1-10l-21-33.3V480c0 17.7-14.3 32-32 32s-32-14.3-32-32v-96h-32z"/>`),
		g.Group(children),
	)
}