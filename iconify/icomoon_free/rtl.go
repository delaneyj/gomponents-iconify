package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 0a4 4 0 0 0 0 8v8h2V2h2v14h2V2h2V0H4zm12 3l-4 4l4 4z"/>`),
		g.Group(children),
	)
}