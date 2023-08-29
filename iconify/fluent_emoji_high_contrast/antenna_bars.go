package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M8.12 27H5.88c-.49 0-.88-.39-.88-.88v-5.25c0-.48.39-.87.88-.87h2.25c.48 0 .87.39.87.88v5.25c0 .48-.39.87-.88.87Zm3.76 0h2.24c.49 0 .88-.39.89-.87V14.88c0-.49-.4-.88-.88-.88h-2.25c-.49 0-.88.39-.88.88v11.24c0 .49.39.88.88.88Zm8.24 0h-2.25c-.48 0-.87-.39-.87-.88V9.88c0-.49.39-.88.88-.88h2.25c.48 0 .87.39.87.88v16.25c0 .48-.39.87-.88.87Zm3.75 0h2.25c.49 0 .88-.39.88-.87V5.88c0-.49-.39-.88-.87-.88h-2.25c-.49 0-.88.39-.88.88v20.24c0 .49.39.88.87.88Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}