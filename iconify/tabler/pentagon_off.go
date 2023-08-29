package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8.133 4.133l2.704-1.965a1.978 1.978 0 0 1 2.326 0l8.021 5.828c.694.504.984 1.397.719 2.212l-1.887 5.808m-.981 3.02l-.196.602a1.978 1.978 0 0 1-1.881 1.367H7.042a1.978 1.978 0 0 1-1.881-1.367l-3.064-9.43a1.978 1.978 0 0 1 .719-2.212L5.81 5.82M3 3l18 18"/>`),
		g.Group(children),
	)
}