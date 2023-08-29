package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotSit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6 19.5h7m-.5-7v-6h-1.172a3 3 0 0 0-2.829 2m5.501 8H6.5v-.25L8 10.4m10 13.1c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4M.5.5l23 23m-11.15-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}