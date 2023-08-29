package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maestro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M8 5.188A7.812 7.812 0 0 0 .187 13A7.812 7.812 0 0 0 13 19a7.812 7.812 0 1 0 0-12a7.772 7.772 0 0 0-5-1.812zm10 2A5.818 5.818 0 0 1 23.813 13A5.818 5.818 0 0 1 18 18.813c-1.378 0-2.628-.51-3.625-1.313a7.79 7.79 0 0 0 1.438-4.5a7.79 7.79 0 0 0-1.438-4.5c.997-.804 2.247-1.313 3.625-1.313z"/>`),
		g.Group(children),
	)
}