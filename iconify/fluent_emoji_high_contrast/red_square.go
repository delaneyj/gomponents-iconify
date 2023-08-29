package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1 28a3 3 0 0 0 3 3h24a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3v24Zm3 1a1 1 0 0 1-1-1v-1.234L5.234 29H4Zm4.062 0L3 23.938v-2.172L10.234 29H8.062Zm5 0L3 18.938v-2.172L15.234 29h-2.172Zm5 0L3 13.938v-2.172L20.234 29h-2.172Zm5 0L3 8.938V6.766L25.234 29h-2.172Zm4.999-.002L3 3.94A1 1 0 0 1 4 3h.234L29 27.766V28a1 1 0 0 1-.94.998Zm.939-4.06L7.062 3h2.172L29 22.766v2.172ZM12.062 3h2.172L29 17.766v2.172L12.062 3Zm5 0h2.172L29 12.766v2.172L17.062 3Zm5 0h2.172L29 7.766v2.172L22.062 3Zm5 0H28a1 1 0 0 1 1 1v.938L27.062 3Z"/>`),
		g.Group(children),
	)
}