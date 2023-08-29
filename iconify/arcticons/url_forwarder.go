package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrlForwarder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.905 20.033H9.093L4.5 27.989h3.812ZM28.182 9.024v10.784h-3.811l-4.594 7.956h8.405v11.212L43.5 23.844Zm-7.75 11.009H16.62l-4.593 7.956h3.811Z"/>`),
		g.Group(children),
	)
}