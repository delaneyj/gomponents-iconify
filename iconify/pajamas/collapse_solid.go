package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0ZM8 8l-.53-.53l-2.25-2.25a.75.75 0 0 1 1.06-1.061l.97.97v-2.38a.75.75 0 0 1 1.5 0v2.38l.97-.97a.75.75 0 1 1 1.06 1.06L8.53 7.47L8 8l.53.53l2.25 2.25a.75.75 0 0 1-1.06 1.061l-.97-.97v2.38a.75.75 0 0 1-1.5 0v-2.38l-.97.97a.75.75 0 1 1-1.06-1.06l2.25-2.25L8 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}