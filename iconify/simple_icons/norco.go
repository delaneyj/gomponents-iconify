package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Norco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.055 2.707a.971.971 0 0 0-.688.387L0 16.78h4.049l7.27-9.597l1.927 5.74l1.42-1.875l-2.578-7.676a.983.983 0 0 0-1.033-.666zM19.95 7.22l-7.27 9.597l-1.927-5.74l-1.42 1.875l2.578 7.676a.987.987 0 0 0 1.72.28L24 7.218z"/>`),
		g.Group(children),
	)
}