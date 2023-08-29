package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsSitOverrides(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m31.707 20.293l-3-3a1 1 0 0 0-1.414 0L18 26.586V31h4.414l9.293-9.293a1 1 0 0 0 0-1.414zm-7.414 6L21.586 29H20v-1.586l2.707-2.707L25 22.414L26.586 24l-2.293 2.293zM28 22.586L26.414 21L28 19.414L29.586 21L28 22.586zM20 20v-2h-4v-7h10v2h2V6c0-1.654-1.346-3-3-3H5C3.346 3 2 4.346 2 6v20c0 1.654 1.346 3 3 3h11v-9h4zm-6-2H4v-7h10v7zM5 5h20a1 1 0 0 1 1 1v3H4V6a1 1 0 0 1 1-1zm9 22H5a1 1 0 0 1-1-1v-6h10v7z"/>`),
		g.Group(children),
	)
}