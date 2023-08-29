package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.71 33.31h1.68c2.78.1 4.49-1.35 5.63-3.62l5.54-11.22c1.38-2.47 3.06-4.14 5.44-4.25h4.39c3.61.08 3.89 2.57 4.12 3.33l2.65 12c.88 2.75 2.43 3.64 4.24 3.7h2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.51 33.48h2.11c2.41 0 4.22-1.41 5.6-4.09l5.09-10.85c1.15-2.43 2.69-4.31 5.69-4.36h2.13M7.54 23.63h15.08"/>`),
		g.Group(children),
	)
}