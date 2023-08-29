package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antivirus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.68 1024h-43l-188-188l-46 46q-14 14-33.5 14t-33.5-14l-410-410q-14-14-14-34t14-34l46-45l-163-163l-46 46q-14 14-33.5 14t-33.5-14l-26-26q-14-14-14-34t14-34l134-134q14-14 34-14t34 14l26 26q14 14 14 33.5t-14 33.5l-46 46l163 162l45-45q14-14 34-14t34 14l410 410q14 14 14 33.5t-14 33.5l-46 46l188 188v43z"/>`),
		g.Group(children),
	)
}