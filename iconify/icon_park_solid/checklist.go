package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChecklist0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m34 10l8 8m0-8l-8 8m10 12l-7 8l-4-4"/><path fill="#fff" d="M26 10H4v8h22v-8Zm0 20H4v8h22v-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChecklist0)"/>`),
		g.Group(children),
	)
}