package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m13.483 380l274-343c26-31 66-31 91 0l274 343c26 32 14 57-27 57h-585c-40 0-53-25-27-57zm37 308h565c28 0 51-22 51-51v-67c0-28-23-50-51-50h-565c-28 0-50 22-50 50v67c0 29 22 51 50 51z"/>`),
		g.Group(children),
	)
}