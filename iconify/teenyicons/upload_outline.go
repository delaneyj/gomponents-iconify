package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m7.5 1.5l3.25 3m-3.25-3l-3 3m3-3V11m6-4v6.5h-12V7"/>`),
		g.Group(children),
	)
}