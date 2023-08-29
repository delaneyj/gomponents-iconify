package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deedum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.803 7.66h2.695v4.552m.88 2.912h2.925v20.911h-3.805V43.5H11.712v-2.85"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.483 4.5v7.465h3.805v20.911h-3.805v7.465H8.697V4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.672 15.644v17.157h8.048V15.644Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.657 12.039h11.064v20.762H16.657z"/>`),
		g.Group(children),
	)
}