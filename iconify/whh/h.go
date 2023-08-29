package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func H(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704.5 1024q-26.5 0-45-18.5T641 960V608q0-13-9.5-22.5T609 576H160q-13 0-22.5 9.5T128 608v352q0 27-18.5 45.5T64 1024t-45.5-18.5T0 960V64q0-26 18.5-45T64 0t45.5 19T128 64v352q0 13 9.5 22.5T160 448h449q13 0 22.5-9.5T641 416V64q0-26 18.5-45t45-19T750 18.5T769 64v896q0 27-19 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}