package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 1024H128q-53 0-90.5-37.5T0 896V128q0-53 37.5-90.5T128 0h512q53 0 90.5 37.5T768 128v768q0 53-37.5 90.5T640 1024zm-255.5-64q26.5 0 45-18.5t18.5-45t-18.5-45.5t-45-19t-45.5 19t-19 45.5t19 45t45.5 18.5zM704 128H64v640h640V128zM128 704V191h384z"/>`),
		g.Group(children),
	)
}