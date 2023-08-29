package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanePrivate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28.584 14.585l-3.12-1.872A4.995 4.995 0 0 0 22.892 12H7.781L7.16 9.515A1.998 1.998 0 0 0 5.22 8H4a2.002 2.002 0 0 0-2 2v7a3.003 3.003 0 0 0 3 3h7v6a2.002 2.002 0 0 0 2 2h1.307a2.009 2.009 0 0 0 1.873-1.298L19.693 20h7.392a2.915 2.915 0 0 0 1.5-5.415ZM27.084 18h-8.777l-3 8H14v-8H5a1 1 0 0 1-1-1v-7h1.22l1 4H10v2h2v-2h3v2h2v-2h3v2h2v-2h.892a2.998 2.998 0 0 1 1.543.428l3.12 1.872a.915.915 0 0 1-.47 1.7Z"/><path fill="currentColor" d="M14 4h1.323l2.4 6h2.154L17.18 3.257A1.99 1.99 0 0 0 15.323 2H14a2.002 2.002 0 0 0-2 2v6h2Z"/>`),
		g.Group(children),
	)
}