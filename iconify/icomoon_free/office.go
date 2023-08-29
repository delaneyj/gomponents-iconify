package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Office(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 16h8V0H0v16zM5 2h2v2H5V2zm0 4h2v2H5V6zm0 4h2v2H5v-2zM1 2h2v2H1V2zm0 4h2v2H1V6zm0 4h2v2H1v-2zm8-5h7v1H9zm0 11h2v-4h3v4h2V7H9z"/>`),
		g.Group(children),
	)
}