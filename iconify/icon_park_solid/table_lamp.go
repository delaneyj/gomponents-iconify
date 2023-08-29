package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTableLamp0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m31 44l11-16l-20-16"/><path fill="#fff" d="m6 12l10 10l8-14l-4-4l-14 8Z"/><path d="M38 44H12m5 0v-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTableLamp0)"/>`),
		g.Group(children),
	)
}