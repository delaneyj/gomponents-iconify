package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M302.41 0h164l558 558v164zm420 0l302 302v164l-466-466h164zm174 0q53 0 90.5 37.5t37.5 90.5v82l-210-210h82zm0 1024h-82L.41 210v-82q0-34 18-64l942 942q-30 18-64 18zm-338 0L.41 466V302l722 722h-164zm-256 0L.41 722V558l466 466h-164zm-174 0q-53 0-90.5-37.5T.41 896v-82l210 210h-82zm0-1024h82l814 814v82q0 34-18 64l-942-942q30-18 64-18z"/>`),
		g.Group(children),
	)
}