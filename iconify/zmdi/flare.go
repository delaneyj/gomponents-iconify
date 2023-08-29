package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 219v42H0v-42h128zm46-69l-30 30l-45-46l30-30zM256 5v128h-43V5h43zm114 129l-45 46l-30-30l45-46zm-29 85h128v42H341v-42zm-106.5-43q26.5 0 45.5 18.5t19 45.5t-19 45.5t-45.5 18.5t-45-18.5T171 240t18.5-45.5t45-18.5zM295 330l30-30l45 46l-30 30zM99 346l45-46l30 30l-45 46zm114 129V347h43v128h-43z"/>`),
		g.Group(children),
	)
}