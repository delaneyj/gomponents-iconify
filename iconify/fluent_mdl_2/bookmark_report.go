package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1664v-640h384v640H0zM1920 384v384h-384V384h384zM1408 0v768h-256v896h-128V0h384zM512 1664V384h384v1280H512zm768-768h768v1152l-384-256l-384 256V896zm640 128h-512v785q65-43 128-85t128-86q65 42 128 85t128 86v-785z"/>`),
		g.Group(children),
	)
}