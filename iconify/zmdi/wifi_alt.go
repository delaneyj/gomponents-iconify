package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128q64-64 149.5-86.5t171 0T469 128l-42 43q-80-80-192.5-80T43 171zm171 171q26-27 63.5-27t64.5 27l-64 64zm-86-86q62-61 149.5-61T384 213l-43 43q-44-44-106.5-44T128 256z"/>`),
		g.Group(children),
	)
}