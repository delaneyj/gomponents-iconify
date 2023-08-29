package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModernMic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.228 10.891a.528.528 0 0 0-.159.69l1.296 2.244c.133.23.438.325.677.208L7 12.116V19h2v-7.854l4.071-1.973l-2.62-4.54l-9.223 6.258zm17.229-7.854a4.061 4.061 0 0 0-5.546-1.484c-.91.525-1.508 1.359-1.801 2.289l2.976 5.156c.951.212 1.973.11 2.885-.415a4.06 4.06 0 0 0 1.486-5.546z"/>`),
		g.Group(children),
	)
}