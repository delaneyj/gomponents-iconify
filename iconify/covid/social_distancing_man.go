package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 5.8A2.4 2.4 0 1 0 12 1a2.4 2.4 0 0 0 0 4.8Zm4.2 5.4a4.2 4.2 0 1 0-8.4 0V13h1.8l.6 6h3.6l.6-6h1.8v-1.8ZM1 6v4m4-2H1m22-2v4m-4-2h4"/><path d="M5 15.914c-2.443.734-4 1.844-4 3.086c0 2.209 4.925 4 11 4s11-1.791 11-4c0-1.242-1.557-2.352-4-3.086"/></g>`),
		g.Group(children),
	)
}