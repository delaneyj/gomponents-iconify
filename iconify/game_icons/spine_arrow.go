package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpineArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M14.563 15.406v104.75L72.97 135.72l52.436 52.468l34.72-18.594l8.81 16.47l-86.186 46.155l119.875 33.155l84.28 84.313l20.22-7.594l6.563 17.5l-65.438 24.594L491 494L381.812 252.72l-26.03 68.31l-17.47-6.655l8.938-23.47l-21.406-21.436l-33.28-120.314l-2.72 5.063l-41.625 77.843l-16.5-8.813l16.686-31.22l-46.375-46.374l-33.06-119.53l-43.064 80.78l-16.5-8.78l15.688-29.407L71.78 15.405H14.563z"/>`),
		g.Group(children),
	)
}