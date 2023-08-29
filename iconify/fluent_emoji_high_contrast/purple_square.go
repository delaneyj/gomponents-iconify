package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PurpleSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31 28a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h24a3 3 0 0 1 3 3v24Zm-3 1a1 1 0 0 0 1-1v-1.234L26.766 29H28Zm-4.062 0L29 23.938v-2.172L21.766 29h2.172Zm-5 0L29 18.938v-2.172L16.766 29h2.172Zm-5 0L29 13.938v-2.172L11.766 29h2.172Zm-5 0L29 8.938V6.766L6.766 29h2.172Zm-4.999-.002L29 3.94A1 1 0 0 0 28 3h-.234L3 27.766V28a1 1 0 0 0 .94.998ZM3 24.938L24.938 3h-2.172L3 22.766v2.172ZM19.938 3h-2.172L3 17.766v2.172L19.938 3Zm-5 0h-2.172L3 12.766v2.172L14.938 3Zm-5 0H7.766L3 7.766v2.172L9.938 3Zm-5 0H4a1 1 0 0 0-1 1v.938L4.938 3Z"/>`),
		g.Group(children),
	)
}