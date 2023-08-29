package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm0-18.236h4.822L16 5.5l-4.822 8.264H16zm.655 9.736H26.5l-4.919-8.44l-2.41 4.131l-2.516 4.309zm-3.825-4.309l-2.411-4.131L5.5 23.5h9.845l-2.515-4.309z"/>`),
		g.Group(children),
	)
}