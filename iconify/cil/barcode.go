package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 464h480V56H16ZM48 88h416v344H48Z"/><path fill="currentColor" d="M80 128h32v256H80zm64 0h32v192h-32zm64 0h32v256h-32zm64 0h32v192h-32zm64 0h32v192h-32zm64 0h32v256h-32zM144 352h32v32h-32zm128 0h32v32h-32zm64 0h32v32h-32z"/>`),
		g.Group(children),
	)
}