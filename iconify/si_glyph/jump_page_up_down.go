package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpPageUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.979 12.021v-1.979H8.02v1.979H6.291a.851.851 0 0 0 0 1.255l2.056 2.459a.96.96 0 0 0 1.334.001l2.053-2.46a.85.85 0 0 0 0-1.255H9.979zM8 3.952v2.046h1.979V3.952h1.73c.365-.36.365.055 0-.307L9.675.305a.949.949 0 0 0-1.326 0l-2.098 3.34c-.365.361-.365-.054 0 .307H8zM1 7h16v2H1z"/>`),
		g.Group(children),
	)
}