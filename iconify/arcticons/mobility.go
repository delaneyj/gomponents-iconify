package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.57 24.75c1.07 1.11 1.63 1.43 2.65 1.39l.11-2.5h1.57m-10.63-1.71h5.19c.13-2.21 1.15-3.74 3.8-4l.07 2.29h1.57M7.63 30.59s-2.76.09-2.89-.92a37.25 37.25 0 0 1 .1-9.23c.25-2.06 3.92-2.3 4.91-4.15c.39-.74-.46-2.08.24-2.1c4.28-.16 8.43-.9 11.39-.68a15.91 15.91 0 0 1 7.36 2.32c2.48 1.48 4.49 3.13 6.84 4.39C38.1 21.58 41.15 21.85 43 24c.73.88.72 5.82-.13 6.25a8.49 8.49 0 0 1-4 .45m-.11-.11a3.93 3.93 0 1 1-3.93-3.93a3.93 3.93 0 0 1 3.88 3.94Zm-23.12 0h8.09m-8.15 0a3.93 3.93 0 1 1-3.93-3.93a3.93 3.93 0 0 1 3.88 3.93Z"/>`),
		g.Group(children),
	)
}