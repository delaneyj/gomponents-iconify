package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledHumanFigure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g stroke="#000" stroke-linecap="round" stroke-linejoin="round"><circle cx="36" cy="36" r="26.68" fill="none" stroke-width="4.75"/><circle cx="36" cy="19.355" r="2" stroke-width="1.33"/><path stroke-width="1.33" d="M30.028 39.967H41.97l-.527-12.658c-.097-1.852-1.616-3.31-3.446-3.33h-3.973a3.514 3.514 0 0 0-3.47 3.33z"/><path stroke-width="1.33" d="M32.694 27.976L31.43 53.29c-.037.733.383 1.333.932 1.333a1.383 1.383 0 0 0 1.152-1.324l1.692-14.676c.084-.729.452-1.324.818-1.324s.733.595.82 1.324L38.535 53.3a1.383 1.383 0 0 0 1.152 1.324c.55 0 .969-.6.932-1.333l-1.266-25.315z"/></g>`),
		g.Group(children),
	)
}