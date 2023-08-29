package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1Z" opacity=".16"/><path fill="currentColor" d="m4 16.127l.98.201a1 1 0 0 0-.092-.66l-.888.46ZM7.873 20l.459-.888a1 1 0 0 0-.66-.092l.2.98ZM3 21l-.98-.201a1 1 0 0 0 1.181 1.18L3 21Zm17-9a8 8 0 0 1-8 8v2c5.523 0 10-4.477 10-10h-2ZM4 12a8 8 0 0 1 8-8V2C6.477 2 2 6.477 2 12h2Zm8-8a8 8 0 0 1 8 8h2c0-5.523-4.477-10-10-10v2ZM4.888 15.668A7.961 7.961 0 0 1 4 12H2c0 1.651.4 3.211 1.112 4.586l1.776-.918ZM12 20a7.968 7.968 0 0 1-3.668-.888l-.918 1.776A9.962 9.962 0 0 0 12 22v-2Zm-8.98-4.074l-1 4.873l1.96.402l1-4.873l-1.96-.402Zm.181 6.054l4.873-1l-.402-1.96l-4.873 1l.402 1.96Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.001v6m-3-3h6"/></g>`),
		g.Group(children),
	)
}