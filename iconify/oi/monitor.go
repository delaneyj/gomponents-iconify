package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.34 0A.5.5 0 0 0 0 .5v5a.5.5 0 0 0 .5.5H3v1H2c-.55 0-1 .45-1 1h6c0-.55-.45-1-1-1H5V6h2.5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM1 1h6v4H1V1z"/>`),
		g.Group(children),
	)
}