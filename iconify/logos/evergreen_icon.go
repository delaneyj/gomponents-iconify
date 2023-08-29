package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvergreenIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><circle cx="128" cy="128" r="128" fill="#38A065"/><path fill="#FFF" d="M135.97 176.049v31.88h-15.94v-31.88h15.94ZM128 56.5l55.79 111.579H72.21L128 56.5Z"/></g>`),
		g.Group(children),
	)
}