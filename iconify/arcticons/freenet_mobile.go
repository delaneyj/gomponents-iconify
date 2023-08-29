package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreenetMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.886 26.46c5.986-.774 15.04 1.505 17.848 2.438c4.931 1.262 4.785 3.337 9.355 7.24c1.4 1.194 2.478-1.335 2.408-2.478c-.283-4.602-9.727-7.008-11.412-9.565c-1.444-2.136 3.59-2.517 1.344-4.954c-2.035-1.8-2.628 2.206-3.69 1.394c-3.25-2.507-8.082-.14-7.812-1.183c.33-1.73 9.003-5.92 6.458-7.35c-3.42-2.256-8.262 4.622-9.386 5.545c-.878.288-1.834-.436-2.747-.302c-1.705.084-4.392 1.294-4.452 2.577c-.115 2.246 3.83 1.635 3.73 2.808c-.23 2.857-11.03 1.403-11.03 5.043s6.187-.812 9.335-1.213Z"/>`),
		g.Group(children),
	)
}