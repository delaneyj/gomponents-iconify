package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Earth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/><path d="M4 24h40"/><path d="M24 44c4.418 0 8-8.954 8-20S28.418 4 24 4s-8 8.954-8 20s3.582 20 8 20Z" clip-rule="evenodd"/><path d="M9.858 10.142A19.937 19.937 0 0 0 24 16a19.937 19.937 0 0 0 14.142-5.858m0 27.716A19.937 19.937 0 0 0 24 32a19.937 19.937 0 0 0-14.142 5.858"/></g>`),
		g.Group(children),
	)
}