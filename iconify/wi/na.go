package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Na(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M6.87 18.23h1.94v-3.64h.02l2.05 3.64h1.99v-6.66h-1.94v3.55h-.02l-1.94-3.55h-2.1v6.66zm6.52.15h1.43l2.61-6.97h-1.42l-2.62 6.97zm2.87-.15h2.07l.29-.95h2.12l.28.95h2.13l-2.43-6.66h-2.01l-2.45 6.66zm2.81-2.39l.64-2.04h.03l.6 2.04h-1.27z"/>`),
		g.Group(children),
	)
}