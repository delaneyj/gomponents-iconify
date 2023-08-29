package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.058 1.4l.659 1.949l-.943.303l-.429-1.273l-8.593 2.906l.686 2.027l-.933.289l-.687-2.027l-.526.215c-.11.22-.132.503-.327.853l-2.004.682l.847 2.507l1.794-.645c.377.162.783.268 1.014.368l10.947-3.696l.216.637l1.183-.4l-1.719-5.093l-1.185.398z"/><path d="M8.916 13.09V8.83c.391-.229.385-.699.387-1.187a1.32 1.32 0 0 0-1.312-1.326a1.32 1.32 0 0 0-1.324 1.316c-.002.494-.003.921.396 1.149v4.255c-1.896.102-3.07 1.287-3.07 1.931h7.999c-.001-.643-1.18-1.778-3.076-1.878z"/></g>`),
		g.Group(children),
	)
}