package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConvexHull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M68.525 4.428a3.5 3.5 0 0 0-.206.061L19.405 20.534a3.5 3.5 0 0 0-2.293 2.432L5.207 67.997a3.5 3.5 0 0 0 1.859 4.046l48.138 23.294a3.5 3.5 0 0 0 4.4-1.155l34.679-49.905a3.5 3.5 0 0 0 .076-3.88L72.361 5.93a3.5 3.5 0 0 0-3.836-1.503Zm-.603 7.558l18.691 29.287l-7.914 2.695l1.612 4.732l3.111-1.06L59.3 82.352l.016-2.105l-4.998-.036l-.052 6.893l-41.553-20.107l10.682-40.405Zm6.045 33.594L64.5 48.804l1.612 4.732l9.467-3.224zm-14.199 4.836l-3.567 1.214a2.5 2.5 0 0 0-1.694 2.349l-.045 6.232l5 .037l.032-4.457l1.886-.643zM54.425 65.21l-.072 10l5 .035l.072-9.998z" color="currentColor"/>`),
		g.Group(children),
	)
}