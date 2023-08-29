package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListTreeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 16.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4 .5h9.503a1 1 0 0 1 .117 1.993l-.117.007H11.5a1 1 0 0 1-.116-1.993L11.5 17h9.503H11.5Zm-8-6.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4 .5h13.503a1 1 0 0 1 .117 1.993l-.117.007H7.5a1 1 0 0 1-.116-1.993L7.5 11h13.503H7.5Zm-4-6.492a1.5 1.5 0 1 1 0 2.999a1.5 1.5 0 0 1 0-3ZM7.5 5h13.503a1 1 0 0 1 .117 1.993L21.003 7H7.5a1 1 0 0 1-.116-1.993L7.5 5h13.503H7.5Z"/>`),
		g.Group(children),
	)
}