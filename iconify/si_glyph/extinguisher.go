package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.995 13.914v-6.83a3.002 3.002 0 0 0-2.03-2.959V2.969h2.032v-.941h-2.032v.035h-.905c-.005-.027-.022-.051-.025-.078c.06-.584.616-1.043 1.291-1.043h1.672V0h-1.672c-1.188 0-2.168.826-2.285 1.879L9.034 2H7.425c-1.51.064-3.355 1.002-3.385 4.078l-1.011 4.828h2.937L5.04 6.082C5.073 3.195 7.329 3.006 7.446 3h1.583v1.125s-1.932.541-1.932 2.959v6.809s-.467 1.951 2.66 1.951h.605c2.88 0 2.633-1.93 2.633-1.93z"/>`),
		g.Group(children),
	)
}