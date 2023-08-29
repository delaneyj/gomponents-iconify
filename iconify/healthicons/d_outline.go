package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 12a3 3 0 0 1 3-3h8c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15h-8a3 3 0 0 1-3-3V12Zm3-1a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h8c7.18 0 13-5.82 13-13s-5.82-13-13-13h-8Zm1 3a1 1 0 0 1 1-1h6c6.075 0 11 4.925 11 11s-4.925 11-11 11h-6a1 1 0 0 1-1-1V14Zm2 1v18h5a9 9 0 1 0 0-18h-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}