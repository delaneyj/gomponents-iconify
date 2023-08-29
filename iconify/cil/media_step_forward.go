package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaStepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M328.845 236.582L134.285 92.463A24 24 0 0 0 96 111.749v288.236a23.979 23.979 0 0 0 38.285 19.286l194.56-144.118a24 24 0 0 0 0-38.57ZM128 384.1V127.63l173.119 128.237ZM384 88h32v336h-32z"/>`),
		g.Group(children),
	)
}