package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1536 64v1408q0 26-19 45t-45 19H960q-26 0-45-19t-19-45V64q0-26 19-45t45-19h512q26 0 45 19t19 45zm-896 0v1408q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h512q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}