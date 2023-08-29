package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notificationlog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.276 32.68V21.47A13.27 13.27 0 0 0 27.08 8.57v-.985a3.102 3.102 0 0 0-6.203 0v.997a13.27 13.27 0 0 0-10.153 12.89V32.68L6.52 36.886v1.942h34.96v-1.942Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.327 38.827a4.651 4.651 0 1 0 9.301.043v-.043m-6.5-10.527h0a2.437 2.437 0 0 1-2.43-2.43v-1.58a2.437 2.437 0 0 1 2.43-2.43h0a2.437 2.437 0 0 1 2.43 2.43v1.58a2.437 2.437 0 0 1-2.43 2.43Zm-6.045-9.72v8.505a1.148 1.148 0 0 0 1.215 1.215h.364m14.255-6.44v7.29a2.437 2.437 0 0 1-2.43 2.43h0a2.039 2.039 0 0 1-1.7-.728"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.487 21.86h0a2.437 2.437 0 0 1 2.43 2.43v1.58a2.437 2.437 0 0 1-2.43 2.43h0a2.437 2.437 0 0 1-2.43-2.43v-1.58a2.437 2.437 0 0 1 2.43-2.43Z"/>`),
		g.Group(children),
	)
}