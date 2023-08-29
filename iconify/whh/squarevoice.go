package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarevoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-257-704q0-53-37.5-90.5t-90.5-37.5t-90.5 37.5t-37.5 90.5v128q0 53 37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5V320zm65 128q0 80-56.5 136t-136 56t-135.5-56t-56-136h-64q0 89 54 157.5t138 90.5v72h-32q-14 0-23 9.5t-9 22.5t9 22.5t23 9.5h191q14 0 23-9.5t9-22.5t-9-22.5t-23-9.5h-31v-72q83-22 137.5-90.5t54.5-157.5h-64z"/>`),
		g.Group(children),
	)
}