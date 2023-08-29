package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.507 3.623l-1.024 1.772c-2.91-.879-5.513-.45-6.41 1.105c-1.178 2.04.79 5.653 4.677 7.897c3.888 2.245 8.001 2.142 9.178.103c.898-1.555-.032-4.024-2.248-6.105l1.023-1.772c3.082 2.71 4.462 6.27 2.957 8.877c-1.86 3.222-7.188 3.355-11.91.63C4.03 13.403 1.48 8.721 3.34 5.5c1.505-2.606 5.28-3.191 9.166-1.877Zm3.377-1.85l1.732 1l-5 8.66l-1.732-1l5-8.66ZM6.732 20H17v2H5.018a.998.998 0 0 1-1.015-.922a.995.995 0 0 1 .131-.578l2.25-3.897l1.732 1L6.732 20Z"/>`),
		g.Group(children),
	)
}