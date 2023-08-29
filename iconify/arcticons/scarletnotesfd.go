package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scarletnotesfd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.57 6.45a56.39 56.39 0 0 1 10.86.76A58.9 58.9 0 0 1 40 10.1c.65.29.42.44-1.51 1c-2.5.78-6.66 2.89-6.85 3.48a14.16 14.16 0 0 0 1.74 3.52c1 1.74 2 3.52 2 3.56s-22.35 4.69-22.69 4.69c-.49 0-5-9.66-6.56-14.1l-.68-2l2.35-1a35.33 35.33 0 0 1 8.77-2.59a24.56 24.56 0 0 1 2.95-.2Zm15.86 15.24c.21 0 .64 1.11 2.57 5c1.21 2.52 2.72 6 3.33 7.74l1.17 3.22l-2.57 1.15c-4.58 2.07-7.91 2.7-13.85 2.74a44.25 44.25 0 0 1-8.32-.63c-3.71-.82-9.92-2.7-10.22-3.19c-.11-.14.34-.4 1-.55c1.63-.34 7.3-2.93 7.68-3.48c.15-.26-.56-1.85-1.62-3.67c-1.25-2.07-2.18-3.6-1.88-3.67c.95-.33 21.86-4.65 22.69-4.69Z"/>`),
		g.Group(children),
	)
}