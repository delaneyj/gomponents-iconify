package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Registered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-miterlimit="10" stroke-width="7.752" d="M27.45 49.57V22.44H37.8a6.757 6.757 0 1 1 0 13.516H27.45" clip-rule="evenodd"/><circle cx="36" cy="36" r="26.68" stroke-width="4.74"/><path stroke-miterlimit="10" stroke-width="7.752" d="m38.03 35.95l5.884 13.62" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}