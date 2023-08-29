package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapdrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.77 42.915a21.5 21.5 0 1 1 20.408.028"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M31.506 38.227a16.125 16.125 0 1 0-15.012 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.826 33.582a10.75 10.75 0 1 0-9.652 0"/><circle cx="24" cy="24" r="4.739" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}