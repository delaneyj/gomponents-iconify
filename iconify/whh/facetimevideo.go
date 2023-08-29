package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facetimevideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.004 768v8q0 16-.5 23t-3 16.5t-9.5 13t-19 3.5l-224-149v149q0 26-19 45t-45 19h-576q-27 0-45.5-18.5t-18.5-45.5V448q0-26 18.5-45t45.5-19h576q26 0 45 19t19 45v149l224-149q12 0 19 3.5t9.5 13t3 16.5t.5 23v264zM.004 192.5q0-79.5 56-136T191.504 0t136 56t56.5 136t-56.5 136t-136 56t-135.5-56t-56-135.5zm448 0q0-79.5 56-136T639.504 0t136 56t56.5 136t-56.5 136t-136 56t-135.5-56t-56-135.5z"/>`),
		g.Group(children),
	)
}