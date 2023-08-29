package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.98 4H7.01A3.014 3.014 0 0 0 4 7.01v9.97A3.025 3.025 0 0 0 7.02 20h9.97A3.014 3.014 0 0 0 20 16.99V7.02A3.025 3.025 0 0 0 16.98 4zM10 5h4v4h-4V5zM9 19H7.02A2.023 2.023 0 0 1 5 16.98V15h4v4zm5 0h-4v-4h4v4zm5-2.01A2.012 2.012 0 0 1 16.99 19H15v-4h4v1.99zM19 14H5V7.01A2.01 2.01 0 0 1 7.01 5H9v4.5a.5.5 0 0 0 .5.5H19v4zm0-5h-4V5h1.98A2.023 2.023 0 0 1 19 7.02V9z"/>`),
		g.Group(children),
	)
}