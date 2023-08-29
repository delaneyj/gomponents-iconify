package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFocusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 157.43V98.57a17 17 0 0 0-8.42-14.71l-51.34-29.65a16.55 16.55 0 0 0-16.48 0L68.43 83.86A17 17 0 0 0 60 98.57v58.86a17 17 0 0 0 8.42 14.71l51.34 29.65a16.53 16.53 0 0 0 16.48 0l51.33-29.65a17 17 0 0 0 8.43-14.71Zm-68-80.26L160.59 96L128 114.81L95.41 96Zm-44 40l32 18.48v36.3l-32-18.53Zm56 54.78V135.6l32-18.48v36.3ZM236 48v40a12 12 0 0 1-24 0V60h-28a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12ZM84 208a12 12 0 0 1-12 12H32a12 12 0 0 1-12-12v-40a12 12 0 0 1 24 0v28h28a12 12 0 0 1 12 12Zm152-40v40a12 12 0 0 1-12 12h-40a12 12 0 0 1 0-24h28v-28a12 12 0 0 1 24 0ZM20 88V48a12 12 0 0 1 12-12h40a12 12 0 0 1 0 24H44v28a12 12 0 0 1-24 0Z"/>`),
		g.Group(children),
	)
}