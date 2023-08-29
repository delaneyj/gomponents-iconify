package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Command(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.5 14A2.502 2.502 0 0 1 9 11.5V10H7v1.5C7 12.879 5.878 14 4.5 14S2 12.879 2 11.5S3.122 9 4.5 9H6V7H4.5C3.122 7 2 5.878 2 4.5S3.122 2 4.5 2S7 3.122 7 4.5V6h2V4.5C9 3.122 10.121 2 11.5 2S14 3.122 14 4.5S12.879 7 11.5 7H10v2h1.5c1.379 0 2.5 1.121 2.5 2.5S12.879 14 11.5 14zM10 10v1.5c0 .827.673 1.5 1.5 1.5s1.5-.673 1.5-1.5s-.673-1.5-1.5-1.5H10zm-5.5 0c-.827 0-1.5.673-1.5 1.5S3.673 13 4.5 13S6 12.327 6 11.5V10H4.5zM7 9h2V7H7v2zm3-3h1.5c.827 0 1.5-.673 1.5-1.5S12.327 3 11.5 3S10 3.673 10 4.5V6zM4.5 3C3.673 3 3 3.673 3 4.5S3.673 6 4.5 6H6V4.5C6 3.673 5.327 3 4.5 3z"/>`),
		g.Group(children),
	)
}