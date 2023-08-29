package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.479 2 2 6.477 2 12v40c0 5.523 4.479 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zm5 43.666A8.332 8.332 0 0 1 48.668 54H15.334A8.334 8.334 0 0 1 7 45.666V12.334A8.334 8.334 0 0 1 15.334 4h33.334A8.332 8.332 0 0 1 57 12.334v33.332z"/><path fill="currentColor" d="M38 45h-6.107V21.979c-2.232 2.086-4.863 3.629-7.893 4.629v-5.543c1.594-.521 3.326-1.512 5.195-2.967c1.871-1.455 3.152-3.156 3.848-5.098H38v32"/>`),
		g.Group(children),
	)
}