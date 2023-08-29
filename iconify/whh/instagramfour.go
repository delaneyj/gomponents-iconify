package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagramfour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320.428 512q0-80 56-136t136-56t136 56t56 136t-56 136t-136 56t-136-56t-56-136zm704-64v448q0 53-37.5 90.5t-90.5 37.5h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v192h-256q-45-59-112-93.5t-144-34.5t-144 34.5t-112 93.5h-256v128h198q-6 32-6 64q0 87 43 160.5t116.5 116.5t160.5 43t160.5-43t116.5-116.5t43-160.5q0-32-6-64h198z"/>`),
		g.Group(children),
	)
}