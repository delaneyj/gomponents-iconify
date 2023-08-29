package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiOneBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.4-2.45 5.5-3.725T12 4q3.425 0 6.525 1.275T24 9L12 21Zm-2.975-5.825q.625-.45 1.388-.7t1.587-.25q.825 0 1.588.25t1.387.7L21.1 9.05q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l6.125 6.125Z"/>`),
		g.Group(children),
	)
}