package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifeBuoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M367.35 80.512a208 208 0 0 0-222.64.074l54.33 69.853a120 120 0 0 1 113.893.035l54.416-69.963zm64.064 64.197l-69.853 54.33a120 120 0 0 1-.035 113.893l69.963 54.416a208 208 0 0 0-.074-222.64zm-280.94 54.357L80.513 144.65a208 208 0 0 0 .074 222.64l69.853-54.33a120 120 0 0 1 .035-113.893zM312.962 361.56a120 120 0 0 1-113.893-.035l-54.416 69.963a208 208 0 0 0 222.64-.074l-54.33-69.853z"/>`),
		g.Group(children),
	)
}