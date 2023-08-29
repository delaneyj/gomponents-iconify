package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 384H826q-23 110-111 183t-203 73t-203-73t-111-183H64q-27 0-45.5-19T0 319.5t18.5-45T64 256h134q23-111 111-183.5T512 0t203 72.5T826 256h134q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zM512 128q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56z"/>`),
		g.Group(children),
	)
}