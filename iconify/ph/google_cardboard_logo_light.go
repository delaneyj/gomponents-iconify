package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleCardboardLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50H32a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h60.69a13.94 13.94 0 0 0 9.9-4.1l24-24a2 2 0 0 1 2.82 0l24 24a13.94 13.94 0 0 0 9.9 4.1H224a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14Zm2 142a2 2 0 0 1-2 2h-60.69a2 2 0 0 1-1.41-.58l-24-24a14 14 0 0 0-19.8 0l-24 24a2 2 0 0 1-1.41.59H32a2 2 0 0 1-2-2V64a2 2 0 0 1 2-2h192a2 2 0 0 1 2 2ZM80 98a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm96-48a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}