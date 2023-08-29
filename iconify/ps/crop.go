package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M427 85H128V0H85v85H0v43h384v256H128V171H85v256h427v-43h-85V85zm-43 384h43v43h-43v-43zM213 213h86v86h-86v-86z"/>`),
		g.Group(children),
	)
}