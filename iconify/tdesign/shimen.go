package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shimen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.999 1.996l11 1.1v4.81l-2.932.293L21.075 22h-7.162L15.01 8.713l-.621.06l-.121.011c-.684.066-1.407.135-2.168.212l-.1.01l-2.914-.292l.991 13.287H2.923L3.931 8.199L.999 7.906v-4.81l11-1.1ZM5.922 8.398L5.075 20h2.847L7.064 8.512l-1.142-.114Zm11.11.113L16.085 20h2.838l-.846-11.6l-.435.046l-.014.001l-.596.063ZM3 4.906v1.19l9 .9c.73-.073 1.422-.14 2.076-.203l.12-.011l1.81-.176c.492-.05.95-.098 1.415-.147l.012-.001c.469-.05.945-.1 1.468-.152L21 6.096v-1.19l-9-.9l-9 .9Z"/>`),
		g.Group(children),
	)
}