package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockKeyLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 82h-34V56a46 46 0 0 0-92 0v26H48a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V96a14 14 0 0 0-14-14ZM94 56a34 34 0 0 1 68 0v26H94Zm116 152a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2V96a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2Zm-82-94a26 26 0 0 0-6 51.29V184a6 6 0 0 0 12 0v-18.71a26 26 0 0 0-6-51.29Zm0 40a14 14 0 1 1 14-14a14 14 0 0 1-14 14Z"/>`),
		g.Group(children),
	)
}