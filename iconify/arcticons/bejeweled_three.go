package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BejeweledThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.177 16.37L35.807 9H12.192l-7.37 7.37a1.102 1.102 0 0 0-.057 1.496L23.164 39.32a1.101 1.101 0 0 0 1.672 0l18.398-21.454a1.102 1.102 0 0 0-.057-1.496ZM5 17.15h38.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.8 39L14 17l4-7.8M24.2 39L34 17l-4-7.8"/>`),
		g.Group(children),
	)
}