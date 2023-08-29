package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioNanny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRadioNanny0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" d="M36 42.001V21.688C36 15.313 31.09 10 24 10s-12 5.313-12 11.688V42c0 1.105.895 1.999 2 1.999h20c1.105 0 2-.894 2-1.999Z"/><circle cx="24" cy="23" r="4" fill="#555"/><path d="M18 34h2m6 4h4M12 20V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRadioNanny0)"/>`),
		g.Group(children),
	)
}