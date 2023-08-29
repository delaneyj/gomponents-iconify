package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20 16a4 4 0 1 0-8 0a4 4 0 0 0 8 0Zm-4-2a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path d="M5 1a4 4 0 0 0-4 4v22a4 4 0 0 0 4 4h22a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5ZM3 5a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v6h-2.088A12 12 0 0 0 16 4C9.373 4 4 9.373 4 16s5.373 12 12 12a12 12 0 0 0 10.912-7H29v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5Zm26 15h-7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h7v8Zm-4.338 1A9.996 9.996 0 0 1 16 26c-5.523 0-10-4.477-10-10S10.477 6 16 6c3.7 0 6.933 2.01 8.662 5H22a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2.662Z"/></g>`),
		g.Group(children),
	)
}