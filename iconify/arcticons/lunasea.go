package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lunasea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.803 13.303a10.342 10.342 0 0 1 2.937.418a8.65 8.65 0 0 0-.861-.044a8.608 8.608 0 0 0-2.526.379a30.168 30.168 0 0 0-2.9 1.546a8.62 8.62 0 1 0 14.004 7.548c.022.281.043.564.043.85a10.698 10.698 0 1 1-20.731-3.683a44.922 44.922 0 0 0-5.021 4.694c-4.38 4.887-9.267 9.52-15.248 9.52c7.338-1.128 12.071-21.229 30.303-21.229Z"/>`),
		g.Group(children),
	)
}