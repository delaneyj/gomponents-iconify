package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MhOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#3b5aa3" d="M0 0h511.9v512H0z"/><path fill="#fff" d="m139 1.2l-5.3 88l-23.2-56.1l9 59.7l-35.9-48.2l23.5 55l-47-36.5L96.7 109L43.5 85.4l46.6 35.3l-58-7.7L87 134.7l-86 7.9l86 5.7l-54.5 22.4L90 163l-46.4 34.2l53.8-23.6l-36.7 46.2l46.7-35.4l-23.4 54l37.4-46.8l-10 58.3l23.4-54.5l5.4 86.1l8.2-85.9l20.3 54.9l-7.7-59.1l37.2 46.8l-24.5-54.7l46.7 37.6l-37-47l55.4 23.1l-49.1-35.8l59.8 10l-57.3-22l89-5.5l-89-8.3L251 116l-60.7 7.6l50.2-35l-56.6 22.7l39-47.3l-47.5 37.1l23-56.8l-37 48.3l8-60.3l-22 56.9l-8.2-88z"/><path fill="#e2ae57" d="M0 498.2L512 0v92.7L0 512v-13.8z"/><path fill="#fff" d="m18 512l494-320.8l-.1-101.9L-.1 512h18z"/></g>`),
		g.Group(children),
	)
}