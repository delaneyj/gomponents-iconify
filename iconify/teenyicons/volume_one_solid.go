package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492V1.5ZM10 6v3h1V6h-1Z"/>`),
		g.Group(children),
	)
}