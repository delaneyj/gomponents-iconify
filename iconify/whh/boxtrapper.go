package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boxtrapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H128q-53 0-90.5-37.5T0 896V128q0-53 37.5-90.5T128 0h576q53 0 90.5 37.5T832 128v768q0 53-37.5 90.5T704 1024zm0-576q0-26-19-45t-45-19H512V160q0-13-9.5-22.5T480 128H352q-13 0-22.5 9.5T320 160v224H192q-27 0-45.5 19T128 448v384q0 27 18.5 45.5T192 896h192v32q0 13 9.5 22.5T416 960t22.5-9.5T448 928v-32h192q27 0 45.5-18.5T704 832V448zm-96 384H448V512h192v288q0 13-9.5 22.5T608 832zm-224 0H224q-13 0-22.5-9.5T192 800V512h192v320z"/>`),
		g.Group(children),
	)
}