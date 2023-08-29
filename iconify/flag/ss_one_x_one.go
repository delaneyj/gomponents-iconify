package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SsOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#078930" d="M0 358.4h512V512H0z"/><path fill="#fff" d="M0 153.6h512v204.8H0z"/><path d="M0 0h512v153.6H0z"/><path fill="#da121a" d="M0 179.2h512v153.6H0z"/><path fill="#0f47af" d="m0 0l433 256L0 512z"/><path fill="#fcdd09" d="M209 207.8L64.4 256l144.8 48.1l-89.5-126v155.8z"/>`),
		g.Group(children),
	)
}