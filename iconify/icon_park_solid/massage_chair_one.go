package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageChairOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMassageChairOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15 28V15.652C15 13 18 10 24 10s9 3 9 5.652V28"/><path fill="#fff" d="M12 34v-6h24v6H12Z"/><path d="M20 4h8M8 16v12h32V16M17 43h14m-7-9v9m0-39v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMassageChairOne0)"/>`),
		g.Group(children),
	)
}