package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm0-144a40 40 0 0 0-28.28 68.28l-10.86 27.28A12 12 0 0 0 100 184h56a12 12 0 0 0 11.14-16.44l-10.86-27.28A40 40 0 0 0 128 72Zm11.31 68.9L150.1 168h-44.2l10.79-27.1a8 8 0 0 0-2.74-9.44a24 24 0 1 1 28.1 0a8 8 0 0 0-2.74 9.44Zm13 32.58Z"/>`),
		g.Group(children),
	)
}