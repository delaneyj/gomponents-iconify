package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdThumbMarkIdentificationPasswordTouchIdSecureFingerprintFingerSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.39a5 5 0 0 0 5-5V5.61a5 5 0 0 0-1.27-3.33M2 6.72v1.67A5 5 0 0 0 5.06 13M9.5 1.28a5 5 0 0 0-6.83 1.83a4.91 4.91 0 0 0-.57 1.52"/><path d="M6.48 3.51A2.51 2.51 0 0 1 9.5 6v1.61m-.64 2.1A2.5 2.5 0 0 1 4.5 8V6a2.53 2.53 0 0 1 .2-1M7 6.11v1.67"/></g>`),
		g.Group(children),
	)
}