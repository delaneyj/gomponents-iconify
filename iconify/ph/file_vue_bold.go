package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVueBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m87.3 156l-20 56a12 12 0 0 1-22.6 0l-20-56a12 12 0 0 1 22.6-8l8.7 24.32L64.7 148a12 12 0 0 1 22.6 8ZM212 196h-20v-4h12a12 12 0 0 0 0-24h-12v-4h20a12 12 0 0 0 0-24h-32a12 12 0 0 0-12 12v56a12 12 0 0 0 12 12h32a12 12 0 0 0 0-24Zm-68-56a12 12 0 0 0-12 12v38a6 6 0 0 1-12 0v-38a12 12 0 0 0-24 0v38a30 30 0 0 0 60 0v-38a12 12 0 0 0-12-12ZM36 108V40a20 20 0 0 1 20-20h96a12 12 0 0 1 8.49 3.52l56 56A12 12 0 0 1 220 88v20a12 12 0 0 1-24 0v-4h-48a12 12 0 0 1-12-12V44H60v64a12 12 0 0 1-24 0Zm124-28h23l-23-23Z"/>`),
		g.Group(children),
	)
}