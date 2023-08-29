package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mastercomfig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.479 0 .174 5.304.174 11.826V24h1.337v-6.716C3.486 21.064 7.446 23.65 12 23.65c4.554 0 8.514-2.586 10.49-6.367V24h1.336V11.826h-1.337c0 5.798-4.69 10.489-10.489 10.489a10.484 10.484 0 0 1-10.49-10.49C1.51 6.028 6.203 1.338 12 1.338zm0 3.72a8.107 8.107 0 1 0 0 16.214a8.107 8.107 0 0 0 0-16.215zm0 1.336a6.77 6.77 0 1 1 0 13.538a6.77 6.77 0 0 1 0-13.538z"/>`),
		g.Group(children),
	)
}