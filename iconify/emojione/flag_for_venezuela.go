package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForVenezuela(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffe62e" d="M32 2C18.9 2 7.8 10.4 3.7 22h56.6C56.2 10.4 45.1 2 32 2z"/><path fill="#ed4c5c" d="M32 62c13.1 0 24.2-8.3 28.3-20H3.7C7.8 53.7 18.9 62 32 62z"/><path fill="#2a5f9e" d="M3.7 22C2.6 25.1 2 28.5 2 32s.6 6.9 1.7 10h56.6c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10H3.7z"/><path fill="#fff" d="m34.6 27.1l1 1.1l-.2-1.5l1.3-.6l-1.4-.3l-.2-1.4l-.7 1.3l-1.5-.3l1 1.1l-.7 1.3zm-5.4.1l1.3.6l-.7-1.3l1-1.1l-1.5.3l-.7-1.3l-.2 1.5l-1.5.2l1.4.7l-.1 1.4zm-5 1.8l1.5.2l-1.2-1l.5-1.3l-1.3.7l-1.2-1l.5 1.5l-1.3.7l1.5.1l.4 1.5zm-4 3.5l1.5-.4l-1.5-.5v-1.5l-.9 1.2l-1.5-.5l.9 1.2l-.8 1.1l1.4-.4l1 1.2zM17.9 37l1.2-.9l-1.6.1l-.5-1.4l-.4 1.4H15l1.3.8l-.4 1.4l1.2-.9l1.4.8zm28.2-.2l-.5 1.4l1.3-.9l1.2.9l-.4-1.4L49 36h-1.6l-.4-1.3l-.5 1.3h-1.6zm-2.5-4.5l.1 1.4l.9-1.1l1.4.3l-.9-1.1l.9-1.2l-1.4.5l-1-1.1v1.4l-1.4.5zm-4-3.4l.6 1.4l.4-1.5l1.5-.1l-1.3-.7l.4-1.5l-1.2 1l-1.3-.7l.6 1.3l-1.2 1z"/>`),
		g.Group(children),
	)
}