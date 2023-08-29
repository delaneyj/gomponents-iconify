package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sennheiser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3v18h24V3zm13.209 1.659c-1.428.548-2.799 1.757-3.905 4.182c-.321.703-.925 2.062-1.2 2.67c-2.224 4.882-3.364 5.932-6.72 5.932V4.35H13.15c.184-.011.235.25.06.309zm9.428 1.894V19.65H10.851c-.181.005-.227-.25-.055-.309c1.427-.548 2.798-1.757 3.904-4.182c.321-.703.926-2.062 1.2-2.67c2.22-4.882 3.36-5.932 6.716-5.932z"/>`),
		g.Group(children),
	)
}