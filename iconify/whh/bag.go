package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1025h-64V257h64q27 0 45.5 18.5T1024 321v640q0 26-18.5 45t-45.5 19zm-128 0H192V257h64v-64q0-80 56-136.5T448 0h128q80 0 136 56.5T768 193v64h64v768zM640 193q0-27-19-45.5T576 129H448q-27 0-45.5 18.5T384 193v64h256v-64zM0 961V321q0-27 18.5-45.5T64 257h64v768H64q-27 0-45.5-19T0 961z"/>`),
		g.Group(children),
	)
}