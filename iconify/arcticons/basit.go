package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.194 20.575C26.76 9.987 20.89 7.89 16.8 7.89S4.5 10.79 4.5 24.839c0 12.615 10.624 15.271 13.559 15.271h16.39c5.312 0 9.05-4.857 9.05-11.008S40.74 17.36 35.184 17.36c-4.403 0-5.836 1.608-6.99 3.215Z"/>`),
		g.Group(children),
	)
}