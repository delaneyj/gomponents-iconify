package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioInputStereoMinijack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4V3c0-.55.45-1 1-1s1 .45 1 1v1h-2m2 5V5h-2v4H9v6c0 1.3.84 2.4 2 2.82V22h2v-4.18c1.16-.42 2-1.52 2-2.82V9h-2Z"/>`),
		g.Group(children),
	)
}