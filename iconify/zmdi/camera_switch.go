package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 45q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67l39-42h128l39 42h68zM277 291l75-75l-75-75v54H149v-54l-74 75l74 75v-54h128v54z"/>`),
		g.Group(children),
	)
}