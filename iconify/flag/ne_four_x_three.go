package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#0db02b" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0h640v320H0z"/><path fill="#e05206" d="M0 0h640v160H0z"/><circle cx="320" cy="240" r="68" fill="#e05206"/>`),
		g.Group(children),
	)
}