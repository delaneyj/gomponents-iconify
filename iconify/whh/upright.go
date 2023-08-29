package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 830H832q-26 0-45-18.5T768 766V442l-567 562q-18 19-45 19t-46-19l-91-90Q0 895 0 868.5T19 823l572-567H256q-26 0-45-19t-19-45V64q0-27 19-45.5T256 0h704q27 0 45.5 18.5T1024 64v702q0 27-19 45.5T960 830z"/>`),
		g.Group(children),
	)
}