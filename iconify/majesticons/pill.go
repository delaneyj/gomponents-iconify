package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.929 19.071a5 5 0 0 1 0-7.071l2.828-2.828l7.071 7.07L12 19.073a5 5 0 0 1-7.071 0zm11.314-4.243L19.07 12A5 5 0 0 0 12 4.929L9.172 7.757l7.07 7.071z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}