package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.812 1.237a3 3 0 0 0-4.243 0L1.591 21.215a1 1 0 0 0-.282.56L.074 30.082a1 1 0 0 0 1.136 1.136l8.307-1.236a1 1 0 0 0 .56-.282L30.054 9.723a3 3 0 0 0 0-4.243l-4.242-4.243Zm-5.92 4.506l5.656 5.657l-2.1 2.1l-5.657-5.657l2.1-2.1Zm2.85 8.464L9.346 27.602L3.69 21.945L17.084 8.55l5.657 5.657ZM8.448 28.12l-4.357.648L2.524 27.2l.649-4.358l5.276 5.277Z"/>`),
		g.Group(children),
	)
}