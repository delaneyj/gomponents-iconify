package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Z(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 128L188 896h516q27 0 45.5 19t18.5 45.5t-18.5 45T704 1024H64q-27 0-45.5-18.5T0 960q0-24 32-64l548-768H64q-27 0-45.5-19T0 63.5t18.5-45T64 0h640q27 0 45.5 18.5T768 64q0 26-32 64z"/>`),
		g.Group(children),
	)
}