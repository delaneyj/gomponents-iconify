package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IHealthServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M21.822 62.675h20.424V42.248h20.429V21.809H42.246V1.383H21.822v20.426H1.388v20.439h20.434z"/>`),
		g.Group(children),
	)
}