package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M704 1216V192q0-14-9-23t-23-9H192q-14 0-23 9t-9 23v1024q0 14 9 23t23 9h480q14 0 23-9t9-23zm672-384V192q0-14-9-23t-23-9H864q-14 0-23 9t-9 23v640q0 14 9 23t23 9h480q14 0 23-9t9-23zm160-768v1408q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1408q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}