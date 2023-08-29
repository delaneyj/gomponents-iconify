package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentosFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13.06l4.47 4.471L12 22l-4.47-4.47L12 13.06Zm-8 3.06L7.879 20H4v-3.88Zm16 0V20h-3.88L20 16.12Zm-2.47-8.59L22 12l-4.469 4.47l-4.47-4.47l4.469-4.47Zm-11.06 0L10.94 12l-4.471 4.469L2 12l4.47-4.47ZM12 2l4.469 4.469L12 10.939L7.53 6.47L12 2ZM7.879 4l-3.88 3.879L4 4h3.879ZM20 4v3.879l-3.88-3.88L20 4Z"/>`),
		g.Group(children),
	)
}