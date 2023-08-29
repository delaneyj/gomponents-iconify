package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.594l-.281-.313l-6-6L19.406 3zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20zM21.094 14c-.551 0-1 .45-1 1s.449 1 1 1c.55 0 1-.45 1-1s-.45-1-1-1zM14 15.594l-.719.687l-4 4l1.438 1.438L14 18.437l2.281 2.282l.719.687l.719-.687L19 19.437l2.281 2.282l1.438-1.438l-3-3l-.719-.687l-.719.687L17 18.563l-2.281-2.282z"/>`),
		g.Group(children),
	)
}