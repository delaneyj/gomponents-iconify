package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#D3D3D3" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#83CBFF" d="M14.854 17.146L8 20l-2.56 6.56c.27.272.646.44 1.06.44h8a.5.5 0 0 0 .5-.5v-9a.498.498 0 0 0-.146-.354ZM17.5 27h8a1.5 1.5 0 0 0 1.5-1.5v-8a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5ZM27 14.5v-8c0-.414-.168-.79-.44-1.06L19 8l-1.854 6.854A.5.5 0 0 0 17.5 15h9a.5.5 0 0 0 .5-.5Z"/><path fill="#00A6ED" d="M5 6.5A1.5 1.5 0 0 1 6.5 5h8a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-8Zm0 11a.5.5 0 0 1 .5-.5h9c.138 0 .263.056.354.146l-9.415 9.415A1.496 1.496 0 0 1 5 25.5v-8ZM26.56 5.44A1.496 1.496 0 0 0 25.5 5h-8a.5.5 0 0 0-.5.5v9c0 .138.056.263.146.354l9.415-9.415Z"/></g>`),
		g.Group(children),
	)
}