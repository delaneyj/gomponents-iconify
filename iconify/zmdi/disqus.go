package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disqus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M219 157q26 0 42 13t16 38t-16 38t-42 13h-23V157h23zm2-157q85 0 145.5 61T427 208t-60.5 147T221 416q-75 0-133-49L0 379l34-85q-18-41-18-86q0-86 60-147T221 0zm112 207q0-46-30.5-74T219 105h-78v206h76q54 0 85-29t31-75z"/>`),
		g.Group(children),
	)
}