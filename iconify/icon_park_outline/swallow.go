package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swallow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m8.999 33.313l15.035-26.37C24.958 4.98 26.627 4 29.04 4c3.62 0 5.977 4.986 5.977 4.986L39 9.58c-4.01.065-6.33.872-6.957 2.42c-.94 2.322 2.456 4.73 2.975 8.004c.52 3.273-1.55 8.801-6.529 11.563c-3.319 1.841-7.462 2.319-12.43 1.433l-6.1 11"/><path d="M18.942 27.978c2.58-4.65 4.275-7.635 5.087-8.957c1.219-1.982 6.245-1.384 4.44 3.563c-1.204 3.298-4.38 5.096-9.527 5.394Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}