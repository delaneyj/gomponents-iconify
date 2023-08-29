package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newpipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.1 4.5L40.9 24L17.64 37.42v-7L28.76 24l-15.6-9v25L7.1 43.5Z"/>`),
		g.Group(children),
	)
}