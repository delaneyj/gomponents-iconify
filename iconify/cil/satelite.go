package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satelite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 496h480V16H16Zm448-32h-25.373l-104-104l54.912-54.911L464 379.55ZM48 48h416v286.3l-74.461-74.461L312 337.373l-112-112l-152 152Zm0 374.627l152-152L393.373 464H48Z"/><path fill="currentColor" d="M120 80H80v40a40 40 0 0 0 40-40Zm-40 83.661V196.6A152.468 152.468 0 0 0 196.6 80h-32.939A120.471 120.471 0 0 1 80 163.661Z"/>`),
		g.Group(children),
	)
}