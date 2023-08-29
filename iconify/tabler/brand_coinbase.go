package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCoinbase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.95 22c-4.503 0-8.445-3.04-9.61-7.413c-1.165-4.373.737-8.988 4.638-11.25a9.906 9.906 0 0 1 12.008 1.598l-3.335 3.367a5.185 5.185 0 0 0-7.354.013a5.252 5.252 0 0 0 0 7.393a5.185 5.185 0 0 0 7.354.013L20 19.088A9.887 9.887 0 0 1 12.95 22z"/>`),
		g.Group(children),
	)
}