package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.29 3.25L5.16 6.72L4 5.34l4.14-3.47l1.15 1.38M22 5.35l-1.16 1.38l-4.14-3.48l1.16-1.38L22 5.35M13 4a8 8 0 0 1 8 8a8 8 0 0 1-8 8a8 8 0 0 1-8-8a8 8 0 0 1 8-8m0 2a6 6 0 0 0-6 6a6 6 0 0 0 6 6a6 6 0 0 0 6-6a6 6 0 0 0-6-6m-1 1.5h1.5v4.53l3.22 1.47l-.62 1.36L12 13V7.5M1 14c0-2.5 1.13-4.7 2.91-6.17A9.975 9.975 0 0 0 3 12l.06 1.13L3 14c0 2.28 1.27 4.26 3.14 5.28c1.3 1.22 2.93 2.11 4.75 2.5c-.61.14-1.24.22-1.89.22a8 8 0 0 1-8-8Z"/>`),
		g.Group(children),
	)
}