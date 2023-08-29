package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLoginPasswordLockLoginPadlockPasswordSecureSecurityTextboxType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y="3.5" rx="1"/><circle cx="3.5" cy="7" r=".5"/><circle cx="6.5" cy="7" r=".5"/><path d="M9.5 8H11"/></g>`),
		g.Group(children),
	)
}