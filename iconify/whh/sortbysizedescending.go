package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sortbysizedescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 192H576q-26 0-45-19t-19-45.5t19-45T576 64h384q27 0 45.5 19t18.5 45t-18.5 45t-45.5 19zm-742 823q-11 8-26 8t-26-8L11 808q-11-8-11-20t11-20h118V64q0-26 18.5-45t45-19T238 19t19 45v704h117q10 8 10 20t-10 20zm358-695h256q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19H576q-26 0-45-19t-19-45.5t19-45t45-18.5zm0 256h128q27 0 45.5 18.5t18.5 45t-18.5 45T704 703H576q-26 0-45-18.5t-19-45t19-45t45-18.5zm.5 255q26.5 0 45 19t18.5 45.5t-18.5 45t-45 18.5t-45.5-18.5t-19-45t19-45.5t45.5-19z"/>`),
		g.Group(children),
	)
}