package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MansShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M2 18.7C2 17.21 3.21 16 4.7 16h8.64c.924 0 1.803.32 2.53.885l-1.588 1.587a.69.69 0 1 0 .976.976l1.64-1.64a.694.694 0 0 0 .085-.103a13.3 13.3 0 0 0 1.716.91l-1.437 1.437a.69.69 0 1 0 .976.976l1.64-1.64a.686.686 0 0 0 .156-.24a25.962 25.962 0 0 0 4.956 1.202c2.9.55 5.01 3.09 5.01 6.05v2.81H14.22c0-1.56-1.26-2.82-2.82-2.82H2V18.7Z"/><path d="M20.034 19.148a.69.69 0 0 0-1.132-.736l-.203.204c.442.197.89.374 1.335.532Zm-3.051-1.443A9.264 9.264 0 0 1 16 16.99a4.481 4.481 0 0 0-.13-.105l.052-.053a.69.69 0 0 1 1.061.873ZM2 27.383h7.98v1.837H2v-1.837Z"/></g>`),
		g.Group(children),
	)
}