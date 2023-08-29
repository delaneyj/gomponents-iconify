package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnockedOutFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FFB02E" d="M15.999 29.998c9.334 0 13.999-6.268 13.999-14c0-7.73-4.665-13.998-14-13.998C6.665 2 2 8.268 2 15.999c0 7.731 4.664 13.999 13.999 13.999Z"/><path fill="#402A32" d="M7.293 9.293a1 1 0 0 1 1.414 0l1.793 1.793l1.793-1.793a1 1 0 1 1 1.414 1.414L11.914 12.5l1.793 1.793a1 1 0 0 1-1.414 1.414L10.5 13.914l-1.793 1.793a1 1 0 0 1-1.414-1.414L9.086 12.5l-1.793-1.793a1 1 0 0 1 0-1.414Zm16 0a1 1 0 1 1 1.414 1.414L22.914 12.5l1.793 1.793a1 1 0 0 1-1.414 1.414L21.5 13.914l-1.793 1.793a1 1 0 0 1-1.414-1.414l1.793-1.793l-1.793-1.793a1 1 0 0 1 1.414-1.414l1.793 1.793l1.793-1.793Z"/><path fill="#BB1D80" d="M12 23a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/></g>`),
		g.Group(children),
	)
}