package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneMarkAsUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 13.33L20 10H7z" opacity=".3"/><path fill="currentColor" d="M7 12v7h13v-7l-6.5 3.33z" opacity=".3"/><path fill="currentColor" d="M16.23 7h2.6c-.06-.47-.36-.94-.79-1.17L10.5 2L2.8 5.83c-.48.26-.8.81-.8 1.34V15c0 1.1.9 2 2 2V7.4L10.5 4l5.73 3z"/><path fill="currentColor" d="M20 8H7c-1.1 0-2 .9-2 2v9c0 1.1.9 2 2 2h13c1.1 0 2-.9 2-2v-9c0-1.1-.9-2-2-2zm0 11H7v-7l6.5 3.33L20 12v7zm-6.5-5.67L7 10h13l-6.5 3.33z"/>`),
		g.Group(children),
	)
}