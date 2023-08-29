package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Now(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M128.1 0L0 232.153h256L128.1 0zm-10.876 46.875l93.256 171.943H22.559l94.665-171.943z" fill="#000"/>`),
		g.Group(children),
	)
}