package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m59.2 0l40 40l-24 24l32 31.9L59.4 128l-40-39.9l7.7-56z"/>`),
		g.Group(children),
	)
}