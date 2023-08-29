package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spamalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M724 1024H300L0 724V300L300 0h424l300 300v424zM576 160q0-13-9.5-22.5T544 128h-64q-13 0-22.5 9.5T448 160v448q0 13 9.5 22.5T480 640h64q13 0 22.5-9.5T576 608V160zm0 640q0-13-9.5-22.5T544 768h-64q-13 0-22.5 9.5T448 800v64q0 13 9.5 22.5T480 896h64q13 0 22.5-9.5T576 864v-64z"/>`),
		g.Group(children),
	)
}