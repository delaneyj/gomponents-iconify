package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M16.4 15.4h3.2v5.2h-3.2z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M21 21a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2H2v9a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-9H21Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="m33.71 12.38l-4.09-4.09a1 1 0 0 0-.7-.29h-5V6.05A2 2 0 0 0 22 4h-8.16A1.92 1.92 0 0 0 12 6.05V8H7.08a1 1 0 0 0-.71.29l-4.08 4.09a1 1 0 0 0-.29.71V17h13v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2h13v-3.92a1 1 0 0 0-.29-.7ZM22 8h-8V6h8Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}