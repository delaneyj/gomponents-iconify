package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bmap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A11.5 11.5 0 0 0 12.5 16c0 9 8.78 19.83 11.08 22.49a.71.71 0 0 0 1 .07l.07-.07C26.92 35.82 35.5 25 35.5 16A11.5 11.5 0 0 0 24 4.5Zm0 14.62a3.9 3.9 0 1 1 3.9-3.89a3.9 3.9 0 0 1-3.9 3.89Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.89 34.13c-3.6.9-6 2.48-6 4.27c0 2.82 5.85 5.1 13.08 5.1s13.08-2.28 13.08-5.1c0-1.77-2.31-3.33-5.83-4.24"/>`),
		g.Group(children),
	)
}