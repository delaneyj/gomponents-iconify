package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 30H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2ZM4 8v20h28V8Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M9 25.3a.8.8 0 0 1-.8-.8v-13a.8.8 0 0 1 1.6 0v13a.8.8 0 0 1-.8.8Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M14.92 25.3a.8.8 0 0 1-.8-.8v-13a.8.8 0 0 1 1.6 0v13a.8.8 0 0 1-.8.8Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M21 25.3a.8.8 0 0 1-.8-.8v-13a.8.8 0 0 1 1.6 0v13a.8.8 0 0 1-.8.8Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M27 25.3a.8.8 0 0 1-.8-.8v-13a.8.8 0 0 1 1.6 0v13a.8.8 0 0 1-.8.8Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}