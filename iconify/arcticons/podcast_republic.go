package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodcastRepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.191 7.021a21.5 21.5 0 1 0 0 33.958m3.093-2.939a21.502 21.502 0 0 0 0-28.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.752 22.564l-11.847-6.84a1.658 1.658 0 0 0-2.486 1.436v13.68a1.658 1.658 0 0 0 2.486 1.436l11.847-6.84a1.658 1.658 0 0 0 0-2.872Z"/>`),
		g.Group(children),
	)
}