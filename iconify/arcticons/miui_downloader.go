package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiuiDownloader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.413 5.5h14.398m-2.757.019v12.526L37.513 35.84c2.212 3.765-1.29 6.659-3.624 6.659H11.416c-2.13 0-5.891-1.978-3.527-6.659c1.785-3.534 10.506-17.796 10.506-17.796V5.664"/><path d="M21.666 24.504c.096-.812.847-1.395 1.677-1.303h0c.83.092 1.425.824 1.329 1.635l-.287 2.424m-2.545-4.225l-.46 3.893"/><path d="M24.672 24.836c.096-.811.847-1.394 1.677-1.302h0c.83.091 1.425.824 1.329 1.635l-.287 2.424m.469 7.305l-.746 2.04c-.25.683.082 1.445.74 1.702s1.393-.087 1.642-.77l.746-2.04m-.746 2.04l-.452 1.236m-13.252-8.303l-1.275 4.815m2.736-1.076l-4.197-2.663m10.046 1.498c-.649 0-1.174-.557-1.174-1.245v-.808c0-.688.525-1.245 1.174-1.245s1.174.557 1.174 1.245v.808c0 .688-.526 1.245-1.174 1.245Zm4.338-4.063l.948 3.11m3.933.481l-.28 3.81m-12.619-7.969l-.207 3.73"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.497 14.575l4.258 4.283l4.258-4.283m-4.258 4.283V5.613"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M22.717 37.43a1.518 1.518 0 0 1-1.261 1.736l-.525.082c-.547.084-1.059-.294-1.144-.845s.288-1.066.835-1.15l2.022-.314"/><path d="M19.76 35.748c.307-.42.557-.491 1.257-.6c.795-.123 1.372.146 1.518 1.091l.418 2.696"/></g>`),
		g.Group(children),
	)
}