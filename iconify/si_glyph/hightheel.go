package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hightheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.396 8.909s-.521-.239-.876 1.241c-.271 1.421-.538 1.734-.538 1.734H2.36s.007-.652.007-1.141s-.283-2.062-.987-3.219C.678 6.367.879 4.881 2.764 3.026c.074-.072.442-.211.902.231c.598.573 3.321 3.391 5.629 4.371c2.168.921 3.627-.121 3.627-.121s2.825.631 3.881 2.413c1.055 1.784-2.296 1.908-2.296 1.908s-4.169.24-5.175-.605c-.151-.192-4.123-2.473-4.936-2.314z"/>`),
		g.Group(children),
	)
}