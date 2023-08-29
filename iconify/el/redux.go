package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M363.858 1007.552L259.669 1123.59H0l259.468-293.53l104.39 177.492zm322.77 116.038h-222.51l-66.197-112.614l134.429-149.7l154.278 262.314zM383.214 985.984L20.943 370.06h222.499l168.881 287.129L925.741 76.41H1200z"/>`),
		g.Group(children),
	)
}