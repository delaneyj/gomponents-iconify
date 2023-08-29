package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignhorizontalcenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 960H448q0 26-18.5 45t-45 19t-45.5-19t-19-45H64q-26 0-45-18.5T0 896V640q0-26 19-45t45-19h256V448H192q-26 0-45-18.5T128 384V128q0-26 19-45t45-19h128q0-26 19-45t45.5-19t45 19T448 64h128q27 0 45.5 19t18.5 45v256q0 27-18.5 45.5T576 448H448v128h256q27 0 45.5 19t18.5 45v256q0 27-18.5 45.5T704 960zM480 320q13 0 22.5-9.5T512 288v-64q0-13-9.5-22.5T480 192H288q-13 0-22.5 9.5T256 224v64q0 13 9.5 22.5T288 320h192z"/>`),
		g.Group(children),
	)
}