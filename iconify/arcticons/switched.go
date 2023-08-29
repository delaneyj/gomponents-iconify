package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switched(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="29.549" height="7.811" x="10.723" y="35.689" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.905" ry="3.905"/><rect width="27.808" height="7.811" x="12.458" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.905" ry="3.905"/><rect width="19.409" height="7.811" x="8.988" y="12.311" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.905" ry="3.905"/><rect width="28.502" height="7.811" x="7.729" y="20.121" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.905" ry="3.905"/><rect width="15.845" height="7.811" x="15.375" y="27.932" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.905" ry="3.905"/>`),
		g.Group(children),
	)
}