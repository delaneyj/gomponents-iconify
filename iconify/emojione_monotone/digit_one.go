package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.568 13.432 30 30 30s30-13.432 30-30C62 15.432 48.568 2 32 2zm6 46h-6.107V24.979c-2.232 2.086-4.863 3.629-7.893 4.629v-5.543c1.594-.521 3.326-1.512 5.195-2.967c1.871-1.455 3.152-3.156 3.848-5.098H38v32z"/>`),
		g.Group(children),
	)
}