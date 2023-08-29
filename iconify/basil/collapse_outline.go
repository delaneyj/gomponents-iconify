package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.028 9.964a.75.75 0 0 0-.75-.75h-2.492V6.722a.75.75 0 0 0-1.5 0v3.242c0 .415.335.75.75.75h3.242a.75.75 0 0 0 .75-.75Zm-3.992 8.064a.75.75 0 0 0 .75-.75v-2.493h2.492a.75.75 0 0 0 0-1.5h-3.242a.75.75 0 0 0-.75.75v3.243c0 .414.335.75.75.75Zm-4.072 0a.75.75 0 0 1-.75-.75v-2.493H6.722a.75.75 0 0 1 0-1.5h3.242a.75.75 0 0 1 .75.75v3.243a.75.75 0 0 1-.75.75ZM5.972 9.964a.75.75 0 0 1 .75-.75h2.492V6.722a.75.75 0 0 1 1.5 0v3.242a.75.75 0 0 1-.75.75H6.722a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}