package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swimsuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSwimsuit0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 4v7m20-7v7M12 31h24v5c-3 0-10 3-10 8h-5c0-5-6-8-9-8v-5Z"/><circle cx="14" cy="18" r="7" fill="#555"/><circle cx="34" cy="18" r="7" fill="#555"/><path d="M21 18h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSwimsuit0)"/>`),
		g.Group(children),
	)
}