package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#67757F" d="M18 31a1 1 0 0 1-1-1V18a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1z"/><path fill="#67757F" d="M18.001 19a1 1 0 0 1-.582-1.814l7-5a1 1 0 0 1 1.163 1.628l-7 4.999a.992.992 0 0 1-.581.187z"/>`),
		g.Group(children),
	)
}