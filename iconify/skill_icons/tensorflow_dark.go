package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TensorflowDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#FF6F00" d="m124.228 229l-33.605-20.11V90.31L40 120.459l.123-44.914L124.228 26v203Zm7.556-203v203l33.609-20.11v-57.109l25.37 15.114l-.151-39.062l-25.219-14.845V90.31L216 120.459l-.122-44.914L131.784 26Z"/></g>`),
		g.Group(children),
	)
}