package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeCrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m117.938 50.5l2.53 99.906l105.313 58.97L117.938 50.5zm260.906 22.594l-97.438 35.97L302.344 280l76.5-206.906zm115.22 78.75l-133.91 59.936l67.563 119.75l5.655 10.064l-11 3.47l-82.063 25.78l57.438 49.25l23.75 20.375l-31.03-4l-254.22-32.814l-35.406-4.562l33.094-13.375l127.187-51.345L173.5 295.03L19.75 363.907v130.656h474.313V151.844zm-357.783 50.47l-96.53 22.655l246.844 72.343l-150.313-95z"/>`),
		g.Group(children),
	)
}