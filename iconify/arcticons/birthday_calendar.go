package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirthdayCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.53v-4.47m0 0c8.606-2.98 0-10.43 0-10.43s-8.942 7.45 0 10.43Zm6.332 11.548a6.333 6.333 0 0 1-12.667 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.666 26.608a6.333 6.333 0 0 1-6.202 6.333A6.431 6.431 0 0 1 5 26.478V24a4.452 4.452 0 0 1 4.433-4.47H38.53A4.452 4.452 0 0 1 43 23.963v2.515a6.431 6.431 0 0 1-6.464 6.463a6.333 6.333 0 0 1-6.203-6.333"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.392 31.692V41.88a1.49 1.49 0 0 1-1.49 1.49H9.096a1.49 1.49 0 0 1-1.49-1.49h0V31.692"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 34.753l.902 1.83l2.02.293L25.46 38.3l.345 2.01l-1.806-.949l-1.806.95l.345-2.011l-1.46-1.424l2.018-.293Z"/>`),
		g.Group(children),
	)
}