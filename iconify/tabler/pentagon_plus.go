package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.5 21.005H7.042a1.978 1.978 0 0 1-1.881-1.367l-3.064-9.43a1.978 1.978 0 0 1 .719-2.212l8.021-5.828a1.978 1.978 0 0 1 2.326 0l8.021 5.828c.694.504.984 1.397.719 2.212l-.78 2.401M16 19h6m-3-3v6"/>`),
		g.Group(children),
	)
}