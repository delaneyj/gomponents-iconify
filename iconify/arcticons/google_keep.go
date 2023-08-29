package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleKeep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a13.51 13.51 0 0 1 13.41 13.41c0 8-7.32 11.09-7.32 15.24v1.21H17.91v-1.21c0-4.17-7.32-7.21-7.32-15.24A13.51 13.51 0 0 1 24 4.5Zm-6.09 29.86h12.18v4.57H17.91Zm0 4.57h12.18v4.57H17.91Z"/>`),
		g.Group(children),
	)
}