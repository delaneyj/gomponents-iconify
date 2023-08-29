package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M118 36Q85 60 65 96t-22 77H0q2-50 25.5-94T88 6zm286 137q-2-41-22.5-77T328 36l31-30q39 29 62 73t26 94h-43zm-42 11v117l43 43v21H42v-21l43-43V184q0-49 30-86.5T191 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5zM223 429q-17 0-29.5-12.5T181 387h85q0 8-3 16q-9 21-31 25q-4 1-9 1z"/>`),
		g.Group(children),
	)
}