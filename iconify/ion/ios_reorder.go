package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosReorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M80 304h352v16H80z" fill="currentColor"/><path d="M80 248h352v16H80z" fill="currentColor"/><path d="M80 192h352v16H80z" fill="currentColor"/>`),
		g.Group(children),
	)
}