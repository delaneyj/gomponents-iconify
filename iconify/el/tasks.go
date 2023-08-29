package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 131.213v234.375h1200V131.213H0zm752.856 58.009h385.62v118.359h-385.62V189.222zM0 482.849v234.375h1200V482.85L0 482.849zm487.72 58.008h650.757v118.358H487.72V540.857zM0 834.412v234.375h1200V834.412H0zm894.946 58.008h243.529v118.359H894.946V892.42z"/>`),
		g.Group(children),
	)
}