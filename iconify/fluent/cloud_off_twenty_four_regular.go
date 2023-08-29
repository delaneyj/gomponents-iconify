package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l4.633 4.634a5.962 5.962 0 0 0-.773 2.105A4.5 4.5 0 0 0 6.5 19h11c.142 0 .282-.006.42-.02l2.8 2.8a.75.75 0 0 0 1.06-1.06L3.28 2.22ZM16.44 17.5H6.5a3 3 0 1 1 0-6h.256c.4 0 .73-.315.749-.715c.03-.631.19-1.229.453-1.766L16.44 17.5Zm4.06-3c0 .782-.3 1.494-.79 2.029l1.062 1.06a4.5 4.5 0 0 0-2.851-7.57a6.001 6.001 0 0 0-8.93-4.21l1.11 1.11a4.5 4.5 0 0 1 6.394 3.866a.75.75 0 0 0 .75.715h.255a3 3 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}