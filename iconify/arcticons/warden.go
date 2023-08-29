package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.55 26.33h5.29l-2.64-4.66l-2.65 4.66z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.11 10.1L26.32 4.92a5.59 5.59 0 0 0-4.25 0L8.94 10.1A2 2 0 0 0 7.62 12v12.47c0 11.16 11.64 19 16.57 19s16.19-7.87 16.19-19V12a2 2 0 0 0-1.27-1.9Zm-7.18 19a2.07 2.07 0 0 1-1.78 1H18.24a2 2 0 0 1-2-2.06a2.11 2.11 0 0 1 .26-1l5.9-10.46a2.07 2.07 0 0 1 2.82-.77a2.11 2.11 0 0 1 .77.77l6 10.51a2.06 2.06 0 0 1-.06 2.05Z"/>`),
		g.Group(children),
	)
}