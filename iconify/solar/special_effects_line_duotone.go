package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecialEffectsLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z" opacity=".5"/><path fill="currentColor" d="M6 8v-.75a.75.75 0 0 0-.75.75H6Zm4 .75a.75.75 0 0 0 0-1.5v1.5Zm0 4a.75.75 0 0 0 0-1.5v1.5Zm8.6-4.3a.75.75 0 1 0-1.2-.9l1.2.9Zm-7.2 7.1a.75.75 0 1 0 1.2.9l-1.2-.9Zm1.2-8a.75.75 0 1 0-1.2.9l1.2-.9Zm4.8 8.9a.75.75 0 1 0 1.2-.9l-1.2.9ZM5.25 16a.75.75 0 0 0 1.5 0h-1.5ZM6 8.75h4v-1.5H6v1.5Zm0 4h4v-1.5H6v1.5Zm11.4-5.2l-3 4l1.2.9l3-4l-1.2-.9Zm-3 4l-3 4l1.2.9l3-4l-1.2-.9Zm-3-3.1l3 4l1.2-.9l-3-4l-1.2.9Zm3 4l3 4l1.2-.9l-3-4l-1.2.9ZM5.25 8v4h1.5V8h-1.5Zm0 4v4h1.5v-4h-1.5Z"/></g>`),
		g.Group(children),
	)
}