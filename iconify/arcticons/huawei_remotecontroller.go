package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiRemotecontroller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.605" height="31.393" x="13.35" y="12.107" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.577"/><rect width="10.853" height="6.492" x="18.725" y="20.678" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.052"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.325 12.107l6.491-7.404M12.184 4.5l7.1 7.607"/>`),
		g.Group(children),
	)
}