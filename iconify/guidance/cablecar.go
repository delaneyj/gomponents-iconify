package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cablecar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 8.5V4.275M23.5 2s-9.609 3-23 3m21 10.863s-4.398-.75-9.5-.75s-9.5.75-9.5.75m6-7.363v6.722m7-6.722v6.722M14.5 2a92.58 92.58 0 0 1-2 .285m-1 .126c-.646.078-1.313.151-2 .219m12 5.87h-19v14h19v-14Z"/>`),
		g.Group(children),
	)
}