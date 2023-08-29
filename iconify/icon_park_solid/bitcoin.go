package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBitcoin0"><g fill="none"><circle cx="24" cy="24" r="20" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#fff" d="M20 16h7a4 4 0 0 1 0 8h-7v-8Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 16v8h7a4 4 0 0 0 4-4v0a4 4 0 0 0-4-4h-2m-5 0h-4m4 0v-4m0 4h5m0 0v-4"/><path fill="#fff" d="M20 24h9a4 4 0 0 1 0 8h-9v-8Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 32v-8h9a4 4 0 0 1 4 4v0a4 4 0 0 1-4 4h-4m-5 0v4m0-4h-4h9m-5 0h5m0 0v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBitcoin0)"/>`),
		g.Group(children),
	)
}