package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stacks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.488 1024h-512q-27 0-45.5-18.5t-18.5-45.5v-64h-64q-27 0-45.5-19t-18.5-45v-64h-64q-27 0-45.5-19t-18.5-45v-64h-64q-27 0-45.5-19t-18.5-45V64q0-27 18.5-45.5T64.488 0h512q27 0 45.5 18.5t18.5 45.5v64h64q27 0 45.5 18.5t18.5 45.5v64h64q26 0 45 18.5t19 45.5v64h64q27 0 45.5 18.5t18.5 45.5v512q0 26-18.5 45t-45.5 19zm-384-960h-512v512h64V192q0-27 18.5-45.5t45.5-18.5h384V64zm128 128h-512v512h64V320q0-27 18.5-45.5t45.5-18.5h384v-64zm128 128h-512v512h64V448q0-27 18.5-45.5t45.5-18.5h384v-64zm128 128h-512v512h512V448z"/>`),
		g.Group(children),
	)
}