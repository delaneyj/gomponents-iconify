package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPiano0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M4 8h40v16H4z"/><path d="M4 24h40v16H4zm6 0v8m6-8v8m10-8v8m6-8v8m6-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPiano0)"/>`),
		g.Group(children),
	)
}