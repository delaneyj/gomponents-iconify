package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 2.5a.5.5 0 0 0-.79-.407l-7 5a.5.5 0 0 0 0 .814l7 5A.5.5 0 0 0 8 12.5V8.472l6.21 4.435A.5.5 0 0 0 15 12.5v-10a.5.5 0 0 0-.79-.407L8 6.528V2.5Z"/>`),
		g.Group(children),
	)
}