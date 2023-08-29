package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unfold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.53 6.51v-4l-1 1l-.71-.71L7.65 1h.71l1.84 1.83l-.71.7l-1-1v3.98h-.96zm0 2.98v4l-1-1l-.71.71L7.65 15h.71l1.84-1.83l-.71-.7l-1 1V9.49h-.96zM13.73 4L14 5.02l-3.68 2.93L14 10.98L13.73 12h-4.2v-1h3L9.55 8.57H6.54L3.45 11h3.08v1H2.27L2 10.98l3.68-2.92L2 5.02L2.27 4h4.26v1H3.45l3 2.42h3.01L12.53 5h-3V4h4.2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}