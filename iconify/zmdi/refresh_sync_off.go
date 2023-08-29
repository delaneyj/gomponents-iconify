package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshSyncOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M152 71q-5 2-16 8l-31-32q22-14 47-20v44zM0 51l27-27l335 336l-27 27l-50-50q-22 14-48 20v-44q7-3 17-8L81 133q-14 28-14 59q0 53 38 90l47-47v128H24l51-51q-51-50-51-120q0-49 26-90zm366-30l-51 51q51 50 51 120q0 49-26 90l-32-31q15-28 15-59q0-53-38-90l-47 47V21h128z"/>`),
		g.Group(children),
	)
}