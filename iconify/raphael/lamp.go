package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 2.833a7 7 0 0 0-7 7c0 3.86 3.945 4.937 4.223 9.5h5.553c.278-4.563 4.224-5.64 4.224-9.5a7 7 0 0 0-7-7zm0 25.333c1.894 0 2.483-1.027 2.667-1.666h-5.334c.184.64.773 1.666 2.667 1.666zm-2.75-2.668h5.5v-5.164h-5.5v5.164z"/>`),
		g.Group(children),
	)
}