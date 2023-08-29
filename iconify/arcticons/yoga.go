package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.849 11.874c3.388 2.733 7.334 7.295-.289 11.492s-11.434-2.945-10.183-6.872a4.27 4.27 0 0 0-1.742 3.003a3.025 3.025 0 0 1-1.328-2.859a6.044 6.044 0 0 0-1.617 1.621s-2.05-7.8 5.515-6.53c.693-1.328 1.877-2.974 5.486-3.263c.116-1.068.704-2.02 2.503-2.512a5.134 5.134 0 0 0 .385 3.032a8.238 8.238 0 0 1 5.919-2.858c.086 3.09-2.319 7.956-2.319 7.956"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.812 22.618c-2.194 1.452-2.31 3.837-5.054 5.137c-3.091 1.464-5.082 1.212-6.872 4.562c-2.079.52-2.021-1.213-4.793-1.155s-5.37 3.234-6.987 5.428s-2.484 3.754-.29 5.197s4.563-3.638 6.93-3.696s4.216 1.04 9.587 1.04a22.22 22.22 0 0 0 10.227-2.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.675 28.967c-2.534-1.068-4.973-4.504-4.424-6.525c-1.068 1.501-1.068 4.418-1.068 4.418c-1.415-1.184-1.386-2.08-1.386-2.08c-1.01 1.79.414 6.823.414 6.823m17.17-.287a12.468 12.468 0 0 0 1.617 8.393c1.348 2.176 2.618 2.336 3.812 1.732c1.676-.85 1.308-2.618.461-4.235a15.387 15.387 0 0 1-.885-8.354"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.773 40.467a2.408 2.408 0 0 0 1.375-.291c1.677-.85 1.309-2.618.462-4.235a14.14 14.14 0 0 1-1.058-6.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.612 22.737c1.74.532 2.047 2.996 2.047 2.996h-1.386a3.892 3.892 0 0 1 .363 4.081a3.073 3.073 0 0 1-1.698-.962s-.819 2.192-1.744 2.751"/><circle cx="31.89" cy="16.523" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.585 19.471s-.161-.307-.32-.624a4.288 4.288 0 0 0 .202-1.328M34.565 18.6a3.618 3.618 0 0 0 1.7.247m-7.888-2.353c.347-.982 1.348-2.66 2.965-2.401m3.84-1.411a2.63 2.63 0 0 1 3.667-.808"/><circle cx="37.194" cy="13.735" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}