package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FFB02E" d="M15.999 29.998c9.334 0 13.999-6.268 13.999-14c0-7.73-4.665-13.998-14-13.998C6.665 2 2 8.268 2 15.999c0 7.731 4.664 13.999 13.999 13.999Z"/><path fill="#fff" d="M10.5 16a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm11 0a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/><path fill="#402A32" d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path fill="#BB1D80" d="M12 22a4 4 0 0 1 8 0v2a4 4 0 0 1-8 0v-2Z"/></g>`),
		g.Group(children),
	)
}