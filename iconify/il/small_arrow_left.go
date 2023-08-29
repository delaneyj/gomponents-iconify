package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M185 414q-10 0-16-7L7 245q-7-7-7-16t7-17L169 50q7-7 17-7t16 7t7 16t-7 17L56 229l146 146q7 7 7 16t-7 16t-17 7z"/>`),
		g.Group(children),
	)
}