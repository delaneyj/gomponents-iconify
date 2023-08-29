package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5c-3.9 0-7 3.1-7 7a6.96 6.96 0 0 0 3.07 5.813C8.51 19.346 6 22.892 6 27h2c0-4.4 3.6-8 8-8c1.2 0 2.4.3 3.4.8c.3-.6.6-1.2 1.1-1.7c-.199-.1-.404-.179-.607-.266A6.96 6.96 0 0 0 23 12c0-3.9-3.1-7-7-7zm0 2c2.8 0 5 2.2 5 5s-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5zm9 11c-2.2 0-4 1.8-4 4v2h-3v8h14v-8h-3v-2c0-2.2-1.8-4-4-4zm0 2c1.1 0 2 .9 2 2v2h-4v-2c0-1.1.9-2 2-2zm-5 6h10v4H20v-4z"/>`),
		g.Group(children),
	)
}