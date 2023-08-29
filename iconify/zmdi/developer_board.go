package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 128h-43v43h43v42h-43v43h43v43h-43v42q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v42h43v43zm-86 213V43H43v298h298zM85 213h107v86H85v-86zm128-42h86v64h-86v-64zM85 85h107v107H85V85zm128 192h86v128h-86V277z"/>`),
		g.Group(children),
	)
}