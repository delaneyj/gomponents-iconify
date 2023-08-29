package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSoapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.928 7.272H4.89a3.358 3.358 0 0 0-3.358 3.358v9.236a3.358 3.358 0 0 0 3.358 3.358h5.038a3.358 3.358 0 0 0 3.359-3.358V10.63a3.358 3.358 0 0 0-3.359-3.358Z"/><path d="M16.488 14.572a2.992 2.992 0 0 0 3.085-4.962a4.967 4.967 0 0 0-7.215-6.698a2.983 2.983 0 0 0-5.844.863m13.961 18.452a1.994 1.994 0 1 0 0-3.987a1.994 1.994 0 0 0 0 3.987ZM7.514 13.254v3.988M5.52 15.248h3.988"/></g>`),
		g.Group(children),
	)
}