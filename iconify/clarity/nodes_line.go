package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NodesLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M10.5 34.29L2 29.39v-9.81l8.5-4.9l8.5 4.9v9.81ZM4 28.23L10.5 32l6.5-3.77v-7.49L10.5 17L4 20.74Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m25.5 34.29l-8.5-4.9v-9.81l8.5-4.9l8.5 4.9v9.81ZM19 28.23L25.5 32l6.5-3.77v-7.49L25.5 17L19 20.74Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="m18 21.32l-8.5-4.9V6.61l8.5-4.9l8.5 4.9v9.81Zm-6.5-6.06L18 19l6.5-3.75V7.77L18 4l-6.5 3.77Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}