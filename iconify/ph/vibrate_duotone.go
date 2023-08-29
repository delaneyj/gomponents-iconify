package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrateDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 56v144a16 16 0 0 1-16 16H96a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h64a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M160 32H96a24 24 0 0 0-24 24v144a24 24 0 0 0 24 24h64a24 24 0 0 0 24-24V56a24 24 0 0 0-24-24Zm8 168a8 8 0 0 1-8 8H96a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8Zm48-112v80a8 8 0 0 1-16 0V88a8 8 0 0 1 16 0Zm32 16v48a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0ZM56 88v80a8 8 0 0 1-16 0V88a8 8 0 0 1 16 0Zm-32 16v48a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}