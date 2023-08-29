package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockLaminatedLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 82h-34V56a46 46 0 0 0-92 0v26H48a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V96a14 14 0 0 0-14-14ZM46 126h164v20H46Zm0 32h164v20H46ZM94 56a34 34 0 0 1 68 0v26H94ZM48 94h160a2 2 0 0 1 2 2v18H46V96a2 2 0 0 1 2-2Zm160 116H48a2 2 0 0 1-2-2v-18h164v18a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}