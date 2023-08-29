package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Today(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M7 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1zm12 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1zM3 2C1.344 2 0 3.344 0 5v18c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V5c0-1.656-1.344-3-3-3h-2v2a2 2 0 0 1-4 0V2H9v2a2 2 0 0 1-4 0V2H3zM2 9h22v14c0 .551-.449 1-1 1H3c-.551 0-1-.449-1-1V9zm14.906 2.156a.575.575 0 0 0-.375.25l-4.906 7.25l-2.281-2.25a.595.595 0 0 0-.844 0l-.875.844a.627.627 0 0 0 0 .875l3.469 3.469c.195.193.505.343.781.343s.572-.176.75-.437l5.906-8.719a.607.607 0 0 0-.156-.844l-1-.687a.64.64 0 0 0-.469-.094zM902 1469v2h26v-2h-26zm4 5v2h18v-2h-18zm-4 5v2h26v-2h-26zm4 5v2h18v-2h-18zm-4 5v2h26v-2h-26z"/>`),
		g.Group(children),
	)
}