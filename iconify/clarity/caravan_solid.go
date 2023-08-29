package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaravanSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M13.5 30C11 30 9 28 9 25.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M33 24h-2v-7.5c0-.5-.1-1-.4-1.5l-4.2-7.5c-.5-1-1.5-1.5-2.6-1.5H5C3.3 6 2 7.3 2 9v14c0 1.7 1.3 3 3 3h2v-2H5c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1h18.8c.4 0 .7.2.9.5l4.2 7.5c.1.1.1.3.1.5V24h-4V12h-7v8h2v-6h3v10h-3v2h13c.6 0 1-.4 1-1s-.4-1-1-1z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M16 18H7v-6h9v6z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}