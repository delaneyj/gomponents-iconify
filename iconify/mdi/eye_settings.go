package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9a3 3 0 0 0-3 3a3 3 0 0 0 3 3a3 3 0 0 0 3-3a3 3 0 0 0-3-3m0 8a5 5 0 0 1-5-5a5 5 0 0 1 5-5a5 5 0 0 1 5 5a5 5 0 0 1-5 5m0-12.5c-4.86 0-9.22 3-11 7.5c2.39 6.08 9.25 9.06 15.33 6.67c3.05-1.2 5.47-3.61 6.67-6.67c-1.78-4.5-6.14-7.5-11-7.5M7 22h2v2H7v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}