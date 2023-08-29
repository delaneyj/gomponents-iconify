package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 6.5H7A2.5 2.5 0 0 0 4.5 9v5A2.5 2.5 0 0 0 7 16.5h6a2.5 2.5 0 0 0 2.5-2.5v-.024l2.348 1.565a.5.5 0 0 0 .777-.416v-7a.5.5 0 0 0-.777-.416L15.5 9.274V9A2.5 2.5 0 0 0 13 6.5Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M11 4.5H5A2.5 2.5 0 0 0 2.5 7v5A2.5 2.5 0 0 0 5 14.5h6a2.5 2.5 0 0 0 2.5-2.5V7A2.5 2.5 0 0 0 11 4.5ZM3.5 7A1.5 1.5 0 0 1 5 5.5h6A1.5 1.5 0 0 1 12.5 7v5a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m15.727 5.58l-2.976 1.936a.5.5 0 0 0-.228.414l-.027 2.612a.5.5 0 0 0 .227.425l3.004 1.952a.5.5 0 0 0 .773-.419V6a.5.5 0 0 0-.773-.42Zm-.227 6l-2.001-1.301l.021-2.07l1.98-1.287v4.658Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}