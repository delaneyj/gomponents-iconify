package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondWithADot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M11.757 8.929a2 2 0 0 0 0 2.828l3.182 3.182a1.5 1.5 0 0 0 2.122 0l3.182-3.182a2 2 0 0 0 0-2.828L17.06 5.747a1.5 1.5 0 0 0-2.122 0l-3.182 3.182Z"/><path fill="#00A6ED" d="M5.747 14.94a1.5 1.5 0 0 0 0 2.12l3.182 3.183a2 2 0 0 0 2.828 0l3.182-3.182a1.5 1.5 0 0 0 0-2.122l-3.182-3.182a2 2 0 0 0-2.828 0L5.747 14.94Z"/><path fill="#00A6ED" d="M11.757 20.243a2 2 0 0 0 0 2.828l3.182 3.182a1.5 1.5 0 0 0 2.122 0l3.182-3.182a2 2 0 0 0 0-2.828L17.06 17.06a1.5 1.5 0 0 0-2.122 0l-3.182 3.182Z"/><path fill="#00A6ED" d="M17.06 17.06a1.5 1.5 0 0 1 0-2.12l3.183-3.183a2 2 0 0 1 2.828 0l3.182 3.182a1.5 1.5 0 0 1 0 2.122l-3.182 3.182a2 2 0 0 1-2.828 0L17.06 17.06Z"/><path fill="#26C9FC" d="M18 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g>`),
		g.Group(children),
	)
}