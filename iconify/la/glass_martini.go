package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassMartini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v1.34l10 13V26h-5v2h12v-2h-5v-6.66l10-13V5H5zm3.031 2H23.97l-1.54 2H11.92l1.54 2h7.434L16 17.36L8.031 7z"/>`),
		g.Group(children),
	)
}