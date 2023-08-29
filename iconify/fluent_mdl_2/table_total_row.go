package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTotalRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm1920 512V256h-512v384h512zm-640 128H768v384h512V768zM768 640h512V256H768v384zm-128 512V768H128v384h512zm768 0h512V768h-512v384zM640 256H128v384h512V256z"/>`),
		g.Group(children),
	)
}