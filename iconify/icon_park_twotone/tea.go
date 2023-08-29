package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTea0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M11 18.167c0-.092.075-.167.167-.167h23.666c.092 0 .167.075.167.167V30c0 6.627-5.373 12-12 12s-12-5.373-12-12V18.167Z"/><path d="M35 30a6 6 0 0 0 0-12"/><path stroke-linejoin="round" d="M11 8v3m24-3v3M23 5v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTea0)"/>`),
		g.Group(children),
	)
}