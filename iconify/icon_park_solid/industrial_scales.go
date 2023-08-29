package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndustrialScales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIndustrialScales0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M10 32h28l4 8H6l4-8Z"/><path stroke-linecap="round" d="M16 40v4m8-32v20"/><path fill="#fff" d="M17 4h14v8H17z"/><path stroke-linecap="round" d="M32 40v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIndustrialScales0)"/>`),
		g.Group(children),
	)
}