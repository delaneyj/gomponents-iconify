package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ffmpeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.72 17.91V6.5l-.53-.49L9.05 18.52l-1.29-.06L24 1.53l-.33-.95l-11.93 1l-5.75 6.6v-.23l4.7-5.39l-1.38-.77l-9.11.77v2.85l1.91.46v.01l.19-.01l-.56.66v10.6c.609-.126 1.22-.241 1.83-.36L14.12 5.22l.83-.04L0 21.44l9.67.82l1.35-.77l6.82-6.74v2.15l-5.72 5.57l11.26.95l.35-.94v-3.16l-3.29-.18a64.66 64.66 0 0 0 1.28-1.23z"/>`),
		g.Group(children),
	)
}