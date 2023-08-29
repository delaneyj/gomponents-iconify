package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixPoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSixPoints0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 24V12m0 12l-10.5 6.062L24 24Zm0 0l10.5 6.062L24 24Z"/><path fill="#555" d="M14 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm14 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm14-8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0-16a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM28 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSixPoints0)"/>`),
		g.Group(children),
	)
}