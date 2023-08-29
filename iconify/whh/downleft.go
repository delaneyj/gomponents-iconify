package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M64 193h128q26 0 45 18.5t19 45.5v325L822 19q19-19 46-19t46 19l91 90q19 19 19 45.5t-19 45.5L433 768h335q26 0 45 18.5t19 44.5v128q0 27-19 45.5t-45 18.5H64q-27 0-45.5-18.5T0 959V257q0-27 18.5-45.5T64 193z"/>`),
		g.Group(children),
	)
}