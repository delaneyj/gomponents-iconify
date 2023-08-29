package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M7 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1zm12 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1zM3 2C1.344 2 0 3.344 0 5v18c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V5c0-1.656-1.344-3-3-3h-2v2a2 2 0 0 1-4 0V2H9v2a2 2 0 0 1-4 0V2H3zM2 9h22v14c0 .551-.449 1-1 1H3c-.551 0-1-.449-1-1V9zm7 3v2.313h4.813l-3.782 7.656H13.5l3.469-8.438V12H9z"/>`),
		g.Group(children),
	)
}