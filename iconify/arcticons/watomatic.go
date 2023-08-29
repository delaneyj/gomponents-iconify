package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watomatic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.4 12.9l-1.76-2.65"/><circle cx="12.26" cy="8.07" r="2.57" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.3" cy="21.43" r="3.81" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.31 12.9l1.75-2.65"/><circle cx="37.44" cy="8.07" r="2.57" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.41" cy="21.43" r="3.81" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.85 10.65c-10 0-18.17 6.88-18.17 15.36a14.21 14.21 0 0 0 5.24 10.79a10.31 10.31 0 0 1-1.56 3.3a11.42 11.42 0 0 1-2.93 2.4a15.44 15.44 0 0 0 9.24-2.79a20.81 20.81 0 0 0 8.18 1.66C34.89 41.37 43 34.49 43 26s-8.11-15.35-18.15-15.35Z"/>`),
		g.Group(children),
	)
}