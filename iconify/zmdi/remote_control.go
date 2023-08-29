package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M230 192q9 0 15 6.5t6 14.5v256q0 9-6 15.5t-15 6.5H102q-9 0-15-6.5T81 469V213q0-8 6-14.5t15-6.5h128zm-64 128q18 0 30.5-12.5t12.5-30t-12.5-30T166 235t-30.5 12.5t-12.5 30t12.5 30T166 320zM60 129q44-44 106-44t106 44l-31 30q-31-31-75-31t-76 31zM166 0q98 0 166 69l-30 30q-56-56-136-56q-79 0-136 56L0 69Q69 0 166 0z"/>`),
		g.Group(children),
	)
}