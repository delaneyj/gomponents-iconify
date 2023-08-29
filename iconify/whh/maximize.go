package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m992 256l-70-70l-117 116q-17 18-42 18t-42.5-17.5T703 260t18-42l116-117l-69-69q0-26 2.5-29T800 0h160q26 0 45 19t19 45v160q0 27-3 29.5t-29 2.5zM768 512v384q0 53-37.5 90.5T640 1024H128q-53 0-90.5-37.5T0 896V512q0-53 37.5-90.5T128 384h512q53 0 90.5 37.5T768 512zm-128 64q0-26-19-45t-45-19H192q-27 0-45.5 19T128 576v256q0 27 18.5 45.5T192 896h384q26 0 45-18.5t19-45.5V576z"/>`),
		g.Group(children),
	)
}