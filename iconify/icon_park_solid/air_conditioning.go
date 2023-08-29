package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirConditioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAirConditioning0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="20" x="4" y="8" rx="2"/><path fill="#fff" d="M12 20h24v8H12z"/><path d="M32 14h4M24 34v6m-8-4v2m16-2v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAirConditioning0)"/>`),
		g.Group(children),
	)
}