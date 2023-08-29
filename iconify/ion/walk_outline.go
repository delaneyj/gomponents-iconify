package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="m314.21 482.32l-56.77-114.74l-44.89-57.39a72.82 72.82 0 0 1-10.13-37.05V144h15.67a40.22 40.22 0 0 1 40.23 40.22v183.36M127.9 293.05v-74.52S165.16 144 202.42 144M370.1 274.42L304 231M170.53 478.36L224 400"/><circle cx="258.32" cy="69.48" r="37.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/>`),
		g.Group(children),
	)
}