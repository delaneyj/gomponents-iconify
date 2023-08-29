package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifebuoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38.142 38.142c7.81-7.81 7.81-20.474 0-28.284c-7.81-7.81-20.474-7.81-28.284 0c-7.81 7.81-7.81 20.474 0 28.284c7.81 7.81 20.474 7.81 28.284 0Zm-5.657-22.627c4.687 4.686 4.687 12.284 0 16.97c-4.686 4.687-12.284 4.687-16.97 0c-4.687-4.686-4.687-12.284 0-16.97c4.686-4.687 12.284-4.687 16.97 0Z" clip-rule="evenodd"/><path d="m38.142 38.142l-5.657-5.657M9.858 38.142l5.657-5.657M9.858 9.858l5.657 5.657m22.627-5.657l-5.657 5.657m0 16.97c4.687-4.686 4.687-12.284 0-16.97c-4.686-4.687-12.284-4.687-16.97 0c-4.687 4.686-4.687 12.284 0 16.97c4.686 4.687 12.284 4.687 16.97 0Z"/></g>`),
		g.Group(children),
	)
}