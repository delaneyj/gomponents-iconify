package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m124.802 0l51.185 15.452L63.959 405.616l-29.938-8.692L124.801 0zM61.04 460.815L15.452 446.77L0 497.955L45.589 512l15.452-51.185zm288.562-63.89l-29.939 8.691L207.637 15.452L258.822 0l90.78 396.924zM338.035 512l45.589-14.045l-15.452-51.185l-45.589 14.045L338.035 512z"/>`),
		g.Group(children),
	)
}