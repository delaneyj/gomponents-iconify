package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarmclockxtreme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="9.25" height="4.62" x="4.41" y="7.52" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".92" transform="rotate(-44.72 9.028 9.839)"/><rect width="4.62" height="9.25" x="36.66" y="5.21" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".92" transform="rotate(-45.28 38.976 9.838)"/><circle cx="24.1" cy="24.84" r="17.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.88 37.76l-1.18 4.92m23.62-4.92l1.18 4.92"/><circle cx="24.1" cy="24.84" r="1.63" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.1 12.28v10.93m1.64 1.63h7.84"/>`),
		g.Group(children),
	)
}