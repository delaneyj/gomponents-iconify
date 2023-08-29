package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTransport0"><g fill="none" stroke-width="4"><rect width="28" height="18" x="16" y="12" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="3"/><path stroke="#000" stroke-linecap="round" d="M24 18v6m12-6v6"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M36 12V6H24v6m20 24H12a2 2 0 0 1-2-2V11a2 2 0 0 0-2-2H4"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M19 42a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Zm18 0a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTransport0)"/>`),
		g.Group(children),
	)
}