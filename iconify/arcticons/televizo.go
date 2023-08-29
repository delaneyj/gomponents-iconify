package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Televizo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.23 23.934l-15.345-8.859v17.719l15.345-8.86zm-15.345-8.859L8.515 9.088m10.37 23.706v12.092M34.23 23.934l10.4-6.004"/>`),
		g.Group(children),
	)
}