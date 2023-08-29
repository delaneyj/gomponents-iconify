package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 0h15v5H0V0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6v7.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V6H1Zm9 3H5V8h5v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}