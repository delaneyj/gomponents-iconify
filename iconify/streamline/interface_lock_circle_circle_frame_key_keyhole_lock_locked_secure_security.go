package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLockCircleCircleFrameKeyKeyholeLockLockedSecureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="6" r="1.5"/><path d="M7 7.5v2"/></g>`),
		g.Group(children),
	)
}