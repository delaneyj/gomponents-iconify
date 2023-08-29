package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minecraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0h169.847v172.501H0zm340.775 0h170.094v172.501H340.775zm0 172.502H169.847v85.915H84.785V512h85.062v-86.668h170.928V512h85.866V258.417h-85.866z"/>`),
		g.Group(children),
	)
}