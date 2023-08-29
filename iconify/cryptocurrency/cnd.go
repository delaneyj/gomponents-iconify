package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm3.473-20.52l2.667-.612l1.033-1.7l-1.02-1.674l-1.306-.902L18.164 6l-1.655.933l-2.648-.472l-3.636 2.52l-.283 2.941l-1.543.644l.178 2.729l-1.077.858l1.076 3.072l.162.322l1.138 3.163l2.34.828l1.733 1.764l1.994.698l.902-.26l1.431-.165l2.098-.601l3.126-1.765l-.818-3.232l-1.574-.62l-.784.897l-2.205.542l-3.207-.508l-1.035-1.197l.36-1.17l-1.516-2.4l1.275-1.556l.136-2.645l1.826-1.048l1.252-.5l1.7.307z"/>`),
		g.Group(children),
	)
}