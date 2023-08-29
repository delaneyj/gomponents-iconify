package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2ZM9.63 24.23a.79.79 0 0 1-.81.77a.79.79 0 0 1-.82-.77V11.77a.79.79 0 0 1 .82-.77a.79.79 0 0 1 .81.77Zm6 0a.79.79 0 0 1-.82.77a.79.79 0 0 1-.81-.77V11.77a.79.79 0 0 1 .81-.77a.79.79 0 0 1 .82.77Zm6.21 0a.79.79 0 0 1-.82.77a.79.79 0 0 1-.81-.77V11.77a.79.79 0 0 1 .81-.77a.79.79 0 0 1 .82.77Zm6.12 0a.79.79 0 0 1-.82.77a.79.79 0 0 1-.81-.77V11.77a.79.79 0 0 1 .81-.77a.79.79 0 0 1 .82.77Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}