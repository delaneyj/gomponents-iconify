package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 .25a.75.75 0 0 1 .673.418l3.058 6.197l6.839.994a.75.75 0 0 1 .415 1.279l-4.948 4.823l1.168 6.811a.751.751 0 0 1-1.088.791L12 18.347l-6.117 3.216a.75.75 0 0 1-1.088-.79l1.168-6.812l-4.948-4.823a.75.75 0 0 1 .416-1.28l6.838-.993L11.328.668A.75.75 0 0 1 12 .25Zm0 2.445L9.44 7.882a.75.75 0 0 1-.565.41l-5.725.832l4.143 4.038a.748.748 0 0 1 .215.664l-.978 5.702l5.121-2.692a.75.75 0 0 1 .698 0l5.12 2.692l-.977-5.702a.748.748 0 0 1 .215-.664l4.143-4.038l-5.725-.831a.75.75 0 0 1-.565-.41L12 2.694Z"/>`),
		g.Group(children),
	)
}