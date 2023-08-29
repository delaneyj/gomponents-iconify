package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M432 112V96a48.14 48.14 0 0 0-48-48H64a48.14 48.14 0 0 0-48 48v256a48.14 48.14 0 0 0 48 48h16"/><rect width="400" height="336" x="96" y="128" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="45.99" ry="45.99"/><ellipse cx="372.92" cy="219.64" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" rx="30.77" ry="30.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M342.15 372.17L255 285.78a30.93 30.93 0 0 0-42.18-1.21L96 387.64M265.23 464l118.59-117.73a31 31 0 0 1 41.46-1.87L496 402.91"/>`),
		g.Group(children),
	)
}