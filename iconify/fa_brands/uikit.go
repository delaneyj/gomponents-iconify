package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M443.9 128v256L218 512L0 384V169.7l87.6 45.1v117l133.5 75.5l135.8-75.5v-151l-101.1-57.6l87.6-53.1L443.9 128zM308.6 49.1L223.8 0l-88.6 54.8l86 47.3l87.4-53z"/>`),
		g.Group(children),
	)
}