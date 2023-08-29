package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 0h254c82 0 157 26 218 71c90 67 149 174 149 296c0 121-59 228-149 296c-61 45-136 72-218 72H0V0zm71 663h189c161-4 290-135 290-296c0-162-129-293-290-296H71v592z"/>`),
		g.Group(children),
	)
}