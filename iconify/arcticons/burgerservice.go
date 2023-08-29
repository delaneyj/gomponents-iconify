package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Burgerservice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.292 15.761V43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M22.15 29.63a6.935 6.935 0 0 1 0 13.87H10.708V15.761H22.15a6.935 6.935 0 1 1 0 13.87Zm0 .001H10.708"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.708 10.145h26.584M10.708 4.5v5.645M37.292 4.5v5.645M24 4.5v5.645"/>`),
		g.Group(children),
	)
}