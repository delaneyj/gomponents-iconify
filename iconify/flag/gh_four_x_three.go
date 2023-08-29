package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#006b3f" d="M0 0h640v480H0z"/><path fill="#fcd116" d="M0 0h640v320H0z"/><path fill="#ce1126" d="M0 0h640v160H0z"/><path d="m320 160l52 160l-136.1-98.9H404L268 320z"/>`),
		g.Group(children),
	)
}