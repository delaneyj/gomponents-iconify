package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUpload0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUpload1" fill="currentColor"><path id="feUpload2" d="M13 5.828V17h-2V5.828L7.757 9.071L6.343 7.657L12 2l5.657 5.657l-1.414 1.414L13 5.828ZM5 19h14a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z"/></g></g>`),
		g.Group(children),
	)
}