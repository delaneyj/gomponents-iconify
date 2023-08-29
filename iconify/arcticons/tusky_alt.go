package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuskyAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.16 11.73C8.7 13.28 4.5 17.42 4.5 24.46s5.75 11.81 13.91 11.81s16-5.38 18-19.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.65 11.73c4.08 1.48 7.61 4.85 9.4 11.19c2 7.11 5 11.06 8.84 11.06s6.61-2.47 6.61-10.32a18.47 18.47 0 0 0-1.06-6.38"/>`),
		g.Group(children),
	)
}