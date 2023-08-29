package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Helm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M491 235h-45q-6-55-41-98l37-37q14-14 0-30q-16-14-30 0l-37 37q-43-35-98-41V21q0-21-21-21t-21 21v45q-59 6-98 41l-37-37q-14-14-30 0q-14 16 0 30l37 37q-35 43-41 98H21q-21 0-21 21t21 21h45q6 55 41 98l-37 37q-14 14 0 30q8 6 15 6q9 0 15-6l37-37q43 35 98 41v45q0 21 21 21t21-21v-45q55-6 98-41l37 37q6 6 15 6q7 0 15-6q14-16 0-30l-37-37q35-43 41-98h45q21 0 21-21t-21-21zm-88 0h-96l68-69q22 29 28 69zm-57-98l-69 68v-96q40 6 69 28zm-111-28v96l-69-68q29-22 69-28zm-98 57l68 69h-96q6-40 28-69zm-28 111h96l-68 69q-22-29-28-69zm57 98l69-68v96q-40-6-69-28zm111 28v-96l69 68q-29 22-69 28zm98-57l-68-69h96q-6 40-28 69z"/>`),
		g.Group(children),
	)
}