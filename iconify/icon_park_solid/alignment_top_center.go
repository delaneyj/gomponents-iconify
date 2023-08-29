package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentTopCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignmentTopCenter0"><g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M18 12v4m6-4v12m6-12v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignmentTopCenter0)"/>`),
		g.Group(children),
	)
}