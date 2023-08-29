package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M16.687 9.649a1 1 0 0 0-1.374 0l-7.011 6.624c-.658.621-.218 1.727.686 1.727h14.024c.904 0 1.344-1.106.686-1.727L16.687 9.65ZM8 22.5a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-15a.5.5 0 0 0-.5.5v2Z"/></g>`),
		g.Group(children),
	)
}