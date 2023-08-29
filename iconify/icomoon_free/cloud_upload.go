package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.942 6.039a2.5 2.5 0 0 0-3.085-2.955a3 3 0 0 0-5.737.075A4 4 0 1 0 4 11h2v3h4v-3h3.5a2.5 2.5 0 0 0 .442-4.961zM9 10v3H7v-3H4.5L8 6.5l3.5 3.5H9z"/>`),
		g.Group(children),
	)
}