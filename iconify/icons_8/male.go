package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 4v2h7.563l-5.97 5.97A8.921 8.921 0 0 0 13 10c-4.96 0-9 4.04-9 9s4.04 9 9 9s9-4.04 9-9a8.94 8.94 0 0 0-1.97-5.594L26 7.436V15h2V4H17zm-4 8c3.878 0 7 3.122 7 7s-3.122 7-7 7s-7-3.122-7-7s3.122-7 7-7z"/>`),
		g.Group(children),
	)
}