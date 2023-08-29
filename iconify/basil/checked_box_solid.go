package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckedBoxSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0a4 4 0 0 1 2.82 1.667L11.5 13.439l-2.47-2.47a.75.75 0 1 0-1.06 1.061l3 3a.75.75 0 0 0 1.06 0l8.116-8.115c.022.11.04.22.053.333c.37 3.157.37 6.346 0 9.504c-.215 1.836-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.644-3.516-3.48a40.903 40.903 0 0 1 0-9.504c.214-1.837 1.69-3.275 3.516-3.48Z"/>`),
		g.Group(children),
	)
}