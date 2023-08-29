package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M366.156 392H328v-72H184v72h-38.627L9.607 256.235L145.845 120H184v72h144v-72h38.627l135.766 135.765ZM54.863 256.235L152 353.373V288h208v64.9l97.137-97.137L360 158.627V224H152v-64.9Z"/>`),
		g.Group(children),
	)
}