package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighHeeledShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 19.98c0-2.15-1.35-3.57-4-3.57l-.92.03l-.07.007V28.5H3V10.01C3 7.37 4.64 5.12 6.95 4.2c1.21-.48 2.59-.07 3.42.93c3.1 3.75 11.66 14.1 13.86 16.28c3.7 3.7 4.56 4.94 4.56 4.94c.72.94.59 2.11-.42 2.11l-15.62-.05c-1.52 0-2.75-1.23-2.75-2.75v-5.68Z"/>`),
		g.Group(children),
	)
}