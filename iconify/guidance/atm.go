package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17.5 6.5c0 4.038.444 8.062 1.322 12M6.5 6.5c0 4.038-.444 8.062-1.322 12M4 6.5h16m-16 4H.5v-8h23v8H20m-12 11V21a2.5 2.5 0 0 0-2.5-2.5h-.322M16 21.5V21a2.5 2.5 0 0 1 2.5-2.5h.322m-13.644 0a55.143 55.143 0 0 1-.427 1.778l-.251.972v.25h15v-.25l-.25-.972a55.09 55.09 0 0 1-.428-1.778M12 14.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}