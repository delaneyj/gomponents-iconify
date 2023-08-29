package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TyphoonLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.654 1.7l-2.782 2.533a9.137 9.137 0 0 1 3.49 1.973c3.512 3.2 3.512 8.388 0 11.588c-2.592 2.36-6.598 3.862-12.016 4.506l2.782-2.533a9.137 9.137 0 0 1-3.49-1.973c-3.512-3.2-3.533-8.369 0-11.588C8.23 3.846 12.237 2.344 17.655 1.7ZM12 6c-3.866 0-7 2.686-7 6s3.134 6 7 6s7-2.686 7-6s-3.134-6-7-6Zm0 2.3c2.21 0 4 1.657 4 3.7c0 2.044-1.79 3.7-4 3.7S8 14.044 8 12c0-2.043 1.79-3.7 4-3.7Zm0 2c-1.138 0-2 .798-2 1.7c0 .903.862 1.7 2 1.7s2-.797 2-1.7c0-.902-.862-1.7-2-1.7Z"/>`),
		g.Group(children),
	)
}