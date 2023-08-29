package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#3F3F3F"><path d="m26.324 22.812l25.295-5.26l-.07-4.677l-25.444 5.666z"/><circle cx="20.756" cy="51.59" r="5.787"/><circle cx="46.206" cy="46.013" r="5.787"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m51.619 17.552l-25.445 5.671l-.069-4.682l25.444-5.666z"/><circle cx="20.756" cy="51.59" r="5.787"/><path d="m26.105 18.541l.438 33.049"/><circle cx="46.206" cy="46.013" r="5.787"/><path d="m51.555 12.963l.438 33.05"/></g>`),
		g.Group(children),
	)
}