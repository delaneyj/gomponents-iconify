package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Normal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.502 20.242s11.599.844 10.91 4.432s-11.511-.078-11.511-.078"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.42 25.58a6.893 6.893 0 0 0 .628 1.927a4.19 4.19 0 0 0 .884.868a4.853 4.853 0 0 1-2.818 1.051c-1.526-.027-2.727-1.094-4.226-1.581a6.135 6.135 0 0 1-2.531-2.269m8.166 3.849s.238 1.754-.31 2.42c-.52.63-1.618.4-2.302.847c-.34.223-.836.516-.937.949c-.074.312.302.914.302.914l.687 1.917a25.957 25.957 0 0 1 3.362 2.47a7.058 7.058 0 0 1 2.256 3.465"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.475 20.706s.27-5.784-1.817-8.228a7.746 7.746 0 0 0-6.796-2.865a7.784 7.784 0 0 0-6.156 3.658c-1.05 1.919-1.005 4.681-1.047 7.173c-.049 2.94.437 8.132 1.621 10.823c.785 1.784 2.932 1.39 3.696 2.61a4.113 4.113 0 0 1 .28 3.832c-.665 1.555-3.43 2.602-4.315 3.725s-1.107 1.095-1.107 1.095"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.256 37.709a10.559 10.559 0 0 0 3.21.785c1.629.098 3.497-1.93 3.497-1.93m-4.2-18.966a1.983 1.983 0 0 0-.078-2.377a2.372 2.372 0 0 0-2.894-.583c-.76.528-.642 1.626-.402 2.52a1.556 1.556 0 0 0 .931.89a2.074 2.074 0 0 0 2.443-.45Zm5.917-.23c.515-.674.086-2.08-.456-2.732a2.018 2.018 0 0 0-2.689-.096c-.76.528-.665 1.105-.425 1.999a1.557 1.557 0 0 0 .719 1.068c.823.355 2.307.473 2.851-.24Z"/><circle cx="21.45" cy="16.479" r=".75" fill="currentColor"/><circle cx="27.233" cy="16.384" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}