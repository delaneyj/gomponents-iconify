package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhosphorLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 26H64a6 6 0 0 0-6 6v128a78.09 78.09 0 0 0 78 78a6 6 0 0 0 6-6v-66h2a70 70 0 0 0 0-140ZM70 54.91L125.74 154H70Zm60 82.19L74.26 38H130ZM70.28 166H130v59.73A66.1 66.1 0 0 1 70.28 166ZM144 154h-2V38h2a58 58 0 0 1 0 116Z"/>`),
		g.Group(children),
	)
}