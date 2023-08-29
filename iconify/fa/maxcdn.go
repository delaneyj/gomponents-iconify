package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1745 517l-164 763h-334l178-832q13-56-15-88q-27-33-83-33h-169l-204 953H620l204-953H538l-204 953H0l204-953L51 0h1276q101 0 189.5 40.5T1664 154q60 73 81 168.5t0 194.5z"/>`),
		g.Group(children),
	)
}