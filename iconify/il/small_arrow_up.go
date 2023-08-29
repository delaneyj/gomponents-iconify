package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M0 310q0-10 7-16l162-163q7-7 17-7t16 7l162 163q7 7 7 16t-7 16t-16 7t-17-7L185 181L39 326q-7 7-16 7t-16-6.5T0 310z"/>`),
		g.Group(children),
	)
}