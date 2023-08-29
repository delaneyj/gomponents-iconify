package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinLocationRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.475.031c-3.007 0-5.443 2.512-5.443 5.609c0 1.584 5.443 10.275 5.443 10.275s5.441-8.609 5.441-10.275c0-3.097-2.437-5.609-5.441-5.609zm2.556 5.985H5.969V4.954h5.062v1.062z"/>`),
		g.Group(children),
	)
}