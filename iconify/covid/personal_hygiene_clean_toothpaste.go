package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneCleanToothpaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.972 20.25h-7.5v1.5a1.5 1.5 0 0 0 1.5 1.5h4.5a1.5 1.5 0 0 0 1.5-1.5v-1.5Zm0 0h-7.5L1.738 2.478A1.5 1.5 0 0 1 3.221.75h10a1.5 1.5 0 0 1 1.482 1.728L11.972 20.25Zm8.308 3v-11l2-3v-8m-4 1h4m-4 3h4m-4 3h4m-14-2v6m-3-3h6"/>`),
		g.Group(children),
	)
}