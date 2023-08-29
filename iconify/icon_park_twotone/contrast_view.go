package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTContrastView0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 7h17v34H7z"/><path stroke-linecap="round" d="M24 7h4m5 0h2m-2 34h2m6-34v2m0 6v2m0 6v2m0 6v2m0 6v2m-14 0h-3m0-37v40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTContrastView0)"/>`),
		g.Group(children),
	)
}