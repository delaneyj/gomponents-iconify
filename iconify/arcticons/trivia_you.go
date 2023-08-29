package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriviaYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5h0a4.876 4.876 0 0 1 4.878 4.875v29.25A4.876 4.876 0 0 1 24 43.5h0a4.876 4.876 0 0 1-4.878-4.875V9.375A4.876 4.876 0 0 1 24 4.5Zm-18.535.89a.889.889 0 0 1 .89-.89h0a8.867 8.867 0 0 1 8.865 8.861v29.25a.889.889 0 0 1-.89.89h0a8.867 8.867 0 0 1-8.866-8.861Zm27.315 7.97a8.867 8.867 0 0 1 8.866-8.86h0a.889.889 0 0 1 .89.89v29.25a8.867 8.867 0 0 1-8.866 8.86h0a.889.889 0 0 1-.89-.89Z"/>`),
		g.Group(children),
	)
}