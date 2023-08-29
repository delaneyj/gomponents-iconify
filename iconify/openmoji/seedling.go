package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seedling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#B1CC33" d="M51.935 35.872c4.2.928 6.765 5.482 6.765 5.482s-4.245 3.05-8.445 2.123s-6.765-5.483-6.765-5.483s4.247-3.048 8.445-2.122z"/><path fill="#5C9E31" d="M22.362 19.992c4.067 1.4 6.098 6.216 6.098 6.216s-4.564 2.548-8.632 1.149s-6.098-6.216-6.098-6.216s4.566-2.546 8.632-1.149z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36 63.166v-12l-5-22m5 22l5-10M22.362 19.992c4.067 1.4 6.098 6.216 6.098 6.216s-4.564 2.548-8.632 1.149s-6.098-6.216-6.098-6.216s4.566-2.546 8.632-1.149zm29.573 15.88c4.2.928 6.765 5.482 6.765 5.482s-4.245 3.05-8.445 2.123s-6.765-5.483-6.765-5.483s4.247-3.048 8.445-2.122z"/>`),
		g.Group(children),
	)
}