package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 16.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4 .5h13.503a1 1 0 0 1 .117 1.993l-.117.007H7.5a1 1 0 0 1-.116-1.993L7.5 17h13.503H7.5Zm-4-6.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4 .5h13.503a1 1 0 0 1 .117 1.993l-.117.007H7.5a1 1 0 0 1-.116-1.993L7.5 11h13.503H7.5Zm-4-6.492a1.5 1.5 0 1 1 0 2.999a1.5 1.5 0 0 1 0-3ZM7.5 5h13.503a1 1 0 0 1 .117 1.993L21.003 7H7.5a1 1 0 0 1-.116-1.993L7.5 5h13.503H7.5Z"/>`),
		g.Group(children),
	)
}