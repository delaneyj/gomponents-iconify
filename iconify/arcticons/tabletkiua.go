package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tabletkiua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="31.184" cy="24" r="12.316"/><path d="M31.184 36.316V11.684"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.005 34.002c-4.813 3.468-11.567 3.037-15.898-1.294c-4.81-4.81-4.81-12.607 0-17.417c4.331-4.33 11.085-4.762 15.899-1.293m-4.928 12.265L8.108 15.29"/>`),
		g.Group(children),
	)
}