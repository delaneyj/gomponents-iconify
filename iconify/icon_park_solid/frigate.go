package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frigate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFrigate0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M11 12v12l13-5l13 5V12H11Z"/><path d="M19 7v5h10V7a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3Zm-7 37L6 26l18-7l18 7l-6 18"/><path d="M4 42s4.663 2 8 2c5 0 7-2 12-2s7 2 12 2c3.337 0 8-2 8-2M24 19v23"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFrigate0)"/>`),
		g.Group(children),
	)
}