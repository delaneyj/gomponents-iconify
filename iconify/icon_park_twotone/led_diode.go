package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LedDiode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLedDiode0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 24a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v6H6v-6Z"/><path d="M19 30v14m10-14v14"/><path fill="#555" d="M24 4c-7.732 0-14 6.268-14 14v4h28v-4c0-7.732-6.268-14-14-14Z"/><circle cx="24" cy="13" r="3" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLedDiode0)"/>`),
		g.Group(children),
	)
}