package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m18.258 14.29l5.951 64.477L50 85.71l25.79-6.943l5.953-64.477H18.258zM68.96 35.187l-18.983 8.116l-.046.019H68.29l-2.1 24.146l-16.184 4.725l-.029-.009v.009l-16.272-4.812l-1.05-12.16h8.077l.525 6.299l8.646 2.183l.074-.021v.01l8.952-2.521l.613-10.148l-9.565-.029l-17.847-.06l-.612-7.611l18.459-7.688l1.076-.447H30.818l-.963-7.788h39.893l-.788 7.787z"/>`),
		g.Group(children),
	)
}