package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M13 4h4v4h-4v9c0 1-.6 3-3 3s-3-2-3-3s.6-3 3-3s3 2 3 3V4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17V8m0 9c0 1-.6 3-3 3s-3-2-3-3s.6-3 3-3s3 2 3 3zm0-9V4h4v4h-4z"/></g>`),
		g.Group(children),
	)
}