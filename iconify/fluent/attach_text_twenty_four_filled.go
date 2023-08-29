package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachTextTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 3a5 5 0 0 1 4.995 4.783L12 8v11a3 3 0 0 1-5.995.176L6 19V9a1 1 0 0 1 1.993-.117L8 9v10a1 1 0 0 0 1.993.117L10 19V8a3 3 0 0 0-5.995-.176L4 8v9a1 1 0 0 1-1.993.117L2 17V8a5 5 0 0 1 5-5Zm8 14h2a1 1 0 0 1 .117 1.993L17 19h-2a1 1 0 0 1-.117-1.993L15 17h2h-2Zm0-4h5.5a1 1 0 0 1 .117 1.993L20.5 15H15a1 1 0 0 1-.117-1.993L15 13h5.5H15Zm0-4h5.5a1 1 0 0 1 .117 1.993L20.5 11H15a1 1 0 0 1-.117-1.993L15 9h5.5H15Zm0-4h5.5a1 1 0 0 1 .117 1.993L20.5 7H15a1 1 0 0 1-.117-1.993L15 5h5.5H15Z"/>`),
		g.Group(children),
	)
}