package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChip0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="24" height="36" x="12" y="6" fill="#fff" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12H6m6 8H6m6 8H6m6 8H6m36-24h-6m6 8h-6m6 8h-6m6 8h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChip0)"/>`),
		g.Group(children),
	)
}