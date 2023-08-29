package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 11.25h3v3h6.5v-3h3s1-9.5-6.25-9.5s-6.25 9.5-6.25 9.5z"/><circle cx="5.25" cy="7.75" r=".5" fill="currentColor"/><circle cx="10.75" cy="7.75" r=".5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}