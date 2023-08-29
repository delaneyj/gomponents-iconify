package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLastColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm768 640v384h512V768H768zm-128 384V768H128v384h512zm640-512V256H768v384h512zm0 640H768v384h512v-384zM640 256H128v384h512V256zM128 1280v384h512v-384H128z"/>`),
		g.Group(children),
	)
}