package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anidroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.785 7.099c-3.336-1.901-7.026-4.668-7.205-.167a9.855 9.855 0 0 0 .333 4.998c-2.735 6.426-2.029 10.916-1.833 31.569h8.497v-9.662c1.832-3.725 4.924-6.026 8.829-5.54V43.5h8.58c.182-11.306.753-30.34-2.333-32.11c.836-8.285.843-8.026-6.747-4.248c-2.415-1.925-4.81-1.566-8.121-.042Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.618 23.175a10.849 10.849 0 0 1 8.538-3.707c1.356-4.823-7.88-9.286-8.538.333Z"/>`),
		g.Group(children),
	)
}