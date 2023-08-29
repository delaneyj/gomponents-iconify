package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Run(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M226.5 85Q209 85 196 72.5t-13-30t13-30T226.5 0t30 12.5t12.5 30t-12.5 30t-30 12.5zM149 381L0 352l9-43l104 21l34-173l-38 15v73H66V145l111-47q3 0 8.5-1t8.5-1q22 0 36 21l22 34q13 23 37.5 37t53.5 14v43q-71 0-117-53l-13 64l45 42v160h-43V330l-44-42z"/>`),
		g.Group(children),
	)
}