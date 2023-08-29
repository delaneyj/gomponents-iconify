package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Zm-21.357-6.404h6.977m-6.977-10.733h6.977m-6.977 5.367h4.549m-4.549-5.367v10.733"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.803 22.625V11.904l5.366 10.733l5.367-10.717v10.717m-10.733 2.726l10.733 10.733m0-10.733L28.803 36.096M16.757 22.605l4.781-10.701m4.582 10.733l-4.582-10.733m3.049 7.142h-6.24"/>`),
		g.Group(children),
	)
}