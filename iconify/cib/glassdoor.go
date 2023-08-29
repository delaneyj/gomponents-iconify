package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glassdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.859 27.427H4.573A4.571 4.571 0 0 0 9.141 32h13.714a4.572 4.572 0 0 0 4.573-4.573V8.656a.165.165 0 0 0-.167-.167h-4.24a.163.163 0 0 0-.161.167v18.776zm0-27.427a4.571 4.571 0 0 1 4.568 4.573H9.146v18.771a.171.171 0 0 1-.167.167h-4.24a.168.168 0 0 1-.167-.167V4.573A4.571 4.571 0 0 1 9.14 0h13.719z"/>`),
		g.Group(children),
	)
}