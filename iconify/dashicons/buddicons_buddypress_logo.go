package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsBuddypressLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.52 0 10 4.48 10 10s-4.48 10-10 10S0 15.52 0 10S4.48 0 10 0zm0 .5C4.75.5.5 4.75.5 10s4.25 9.5 9.5 9.5s9.5-4.25 9.5-9.5S15.25.5 10 .5zm0 1c4.7 0 8.5 3.8 8.5 8.5s-3.8 8.5-8.5 8.5s-8.5-3.8-8.5-8.5S5.3 1.5 10 1.5zm1.8 1.71c-.57 0-1.1.17-1.55.45a3.55 3.55 0 0 1 2.73 3.45c0 .69-.21 1.33-.55 1.87a2.917 2.917 0 0 0 2.29-2.85c0-1.61-1.31-2.92-2.92-2.92zm-2.38 1a2.926 2.926 0 1 0 .011 5.851A2.926 2.926 0 0 0 9.42 4.21zm4.25 5.01l-.51.59c2.34.69 2.45 3.61 2.45 3.61h1.28c0-4.71-3.22-4.2-3.22-4.2zm-2.1.8l-2.12 2.09l-2.12-2.09C3.12 10.24 3.89 15 3.89 15h11.08c.47-4.98-3.4-4.98-3.4-4.98z"/>`),
		g.Group(children),
	)
}