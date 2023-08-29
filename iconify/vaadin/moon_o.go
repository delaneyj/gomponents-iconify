package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.2 11.9c-4.5 0-8.1-3.6-8.1-8.1c0-1.4.3-2.7.9-3.8c-3.4.9-6 4.1-6 7.9C0 12.4 3.6 16 8.1 16c3.1 0 5.8-1.8 7.2-4.4c-.6.2-1.3.3-2.1.3zM8.1 15C4.2 15 1 11.8 1 7.9c0-2.5 1.3-4.7 3.3-6c-.2.6-.2 1.2-.2 1.9c0 5 4.1 9.1 9.1 9.2c-1.4 1.2-3.2 2-5.1 2z"/>`),
		g.Group(children),
	)
}