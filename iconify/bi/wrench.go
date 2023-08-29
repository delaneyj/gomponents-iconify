package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.102 2.223A3.004 3.004 0 0 0 3.78 5.897l6.341 6.252A3.003 3.003 0 0 0 13 16a3 3 0 1 0-.851-5.878L5.897 3.781A3.004 3.004 0 0 0 2.223.1l2.141 2.142L4 4l-1.757.364L.102 2.223zm13.37 9.019l.528.026l.287.445l.445.287l.026.529L15 13l-.242.471l-.026.529l-.445.287l-.287.445l-.529.026L13 15l-.471-.242l-.529-.026l-.287-.445l-.445-.287l-.026-.529L11 13l.242-.471l.026-.529l.445-.287l.287-.445l.529-.026L13 11l.471.242z"/>`),
		g.Group(children),
	)
}