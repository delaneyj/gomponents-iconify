package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPageTemplate0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M23 4H4v22h19V4Zm21 30H4v9h40v-9Zm0-30H31v8h13V4Zm0 14H31v8h13v-8Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPageTemplate0)"/>`),
		g.Group(children),
	)
}