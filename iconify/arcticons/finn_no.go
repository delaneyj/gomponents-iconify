package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinnNo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.309 33.925H24V15.267c0-.658.533-1.192 1.192-1.192h17.117c.658 0 1.191.534 1.191 1.192v17.468c0 .658-.533 1.192-1.191 1.192Zm-18.309 0H5.691A1.192 1.192 0 0 1 4.5 32.734V15.266c0-.658.533-1.192 1.191-1.192h5.314C18.182 14.074 24 19.892 24 27.07"/>`),
		g.Group(children),
	)
}