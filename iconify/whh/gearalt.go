package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gearalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 576H890q-14 87-66 159l63 64q9 9 9 22t-9 22l-44 44q-9 9-22 9t-22-9l-64-63q-72 52-159 66v102q0 13-9.5 22.5T544 1024h-64q-13 0-22.5-9.5T448 992V890q-87-14-159-66l-64 63q-9 9-22 9t-22-9l-44-44q-9-9-9-22t9-22l63-64q-52-72-66-159H32q-13 0-22.5-9.5T0 544v-64q0-13 9.5-22.5T32 448h102q14-86 66-159l-63-63q-9-9-9-22t9-22l44-45q9-9 22-9t22 9l64 64q72-52 159-67V32q0-13 9.5-22.5T480 0h64q13 0 22.5 9.5T576 32v102q87 15 159 67l64-64q9-9 22-9t22 9l44 45q9 9 9 22t-9 22l-63 63q52 73 66 159h102q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 576zM512 256q-106 0-181.5 75T255 512t75.5 181T512 768t181.5-75T769 512t-75.5-181T512 256z"/>`),
		g.Group(children),
	)
}