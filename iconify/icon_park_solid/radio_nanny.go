package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioNanny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRadioNanny0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M36 42.001V21.688C36 15.313 31.09 10 24 10s-12 5.313-12 11.688V42c0 1.105.895 1.999 2 1.999h20c1.105 0 2-.894 2-1.999Z"/><circle cx="24" cy="23" r="4" fill="#000" stroke="#000"/><path stroke="#000" d="M18 34h2m6 4h4"/><path stroke="#fff" d="M12 20V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRadioNanny0)"/>`),
		g.Group(children),
	)
}