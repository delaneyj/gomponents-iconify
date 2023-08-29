package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoftoffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 1024L0 832l576 64V128L192 256v512L0 832V192L576 0l320 128v768z"/>`),
		g.Group(children),
	)
}