package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backbonedotjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.34 0v10.45l3.2-1.83V5.27l2.93 1.67l3.01-1.72L2.34 0zm19.31 0L12.5 5.22l3.02 1.73l2.94-1.68v3.35l3.2 1.83V0h-.01zm-9.9 5.64l-9.4 5.38V24l9.4-5.36v-3.76l-6.21 3.56v-5.5l6.21-3.54V5.64zm.5 0V9.4l6.22 3.54v5.5l-6.22-3.56v3.76L21.66 24V11.02l-9.41-5.38zM7.7 12.3l-1.65.94v1.86l2.17 1.24l3.28-1.87l-3.8-2.17zm8.61 0l-3.8 2.16l3.28 1.88l2.17-1.24v-1.86l-1.65-.94z"/>`),
		g.Group(children),
	)
}