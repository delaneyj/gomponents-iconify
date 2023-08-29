package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHelpCustomerSupportFiveCustomerHeadsetHelpPhoneSupport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3" height="6" x="2.5" y="7.5" rx="1"/><rect width="3" height="6" x="8.5" y="7.5" rx="1"/><path d="M.5 9.5V7a6.5 6.5 0 0 1 13 0v2.5"/></g>`),
		g.Group(children),
	)
}