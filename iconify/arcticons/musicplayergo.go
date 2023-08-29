package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicplayergo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.76 17.75v14.66a1 1 0 0 0 1.06 1a1.06 1.06 0 0 0 .57-.17L31.2 26a1 1 0 0 0 .29-1.44a1.33 1.33 0 0 0-.29-.29l-10.8-7.39a1 1 0 0 0-1.64.87Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.87 5.75a19.43 19.43 0 1 0 21.56 19.32"/><circle cx="24" cy="5.63" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}