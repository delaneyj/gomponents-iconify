package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM76 128a12 12 0 0 1 12-12h28V88a12 12 0 0 1 24 0v28h28a12 12 0 0 1 0 24h-28v28a12 12 0 0 1-24 0v-28H88a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}