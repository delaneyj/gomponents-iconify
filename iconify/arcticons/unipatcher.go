package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unipatcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.524h5.924S35.172 37 40.479 36.578s2.95-11.169-1.887-15.83S27.79 16.524 24 16.524s-9.754-.435-14.592 4.226s-7.193 15.408-1.886 15.829s10.554-5.054 10.554-5.054H24m-3.846-4.434h2.545m5.147 0h-2.545"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.588 21.287a1.355 1.355 0 0 0-1.359 1.359v1.165h-1.165a1.36 1.36 0 0 0 0 2.72h1.165v1.164a1.358 1.358 0 0 0 2.717 0V26.53h1.168a1.36 1.36 0 0 0 0-2.72h-1.168v-1.164a1.355 1.355 0 0 0-1.358-1.359Z"/><circle cx="36.59" cy="25.128" r="1.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.102" cy="25.128" r="1.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.846" cy="27.872" r="1.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.846" cy="22.384" r="1.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.523s-.79-2.597 1.647-2.545s1.999-2.579 1.999-2.579"/>`),
		g.Group(children),
	)
}