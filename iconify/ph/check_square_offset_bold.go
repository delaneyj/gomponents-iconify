package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquareOffsetBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 48v160a20 20 0 0 1-20 20h-68a12 12 0 0 1 0-24h64V52H52v88a12 12 0 0 1-24 0V48a20 20 0 0 1 20-20h160a20 20 0 0 1 20 20Zm-99.51 103.51a12 12 0 0 0-17 0L64 199l-15.51-15.49a12 12 0 1 0-17 17l24 24a12 12 0 0 0 17 0l56-56a12 12 0 0 0 0-17Z"/>`),
		g.Group(children),
	)
}