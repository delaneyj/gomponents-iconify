package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h2m6 7c-.667 5.333-2.333 8-5 8h-4c-2.667 0-4.333-2.667-5-8"/><path d="M11 16c0 .667.333 1 1 1s1-.333 1-1h-2zm1 2v2m-2-9v.01m4-.01v.01M5 4l6 .97l-6.238 6.688a1.021 1.021 0 0 1-1.41.111a.953.953 0 0 1-.327-.954L5 4zm14 0l-6 .97l6.238 6.688c.358.408.989.458 1.41.111a.953.953 0 0 0 .327-.954L19 4z"/></g>`),
		g.Group(children),
	)
}