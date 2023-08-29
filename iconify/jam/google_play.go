package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.625 9.267l-3.635-2.1L11 10.061l2.733 2.644l3.892-2.25a.68.68 0 0 0 .342-.593a.68.68 0 0 0-.342-.595zM13.373 6.81l-4-2.312L.607 0l9.901 9.584zM.752 19.98l8.636-4.763l3.728-2.155l-2.608-2.524zM.022.388l-.01 19.355l10.003-9.682z"/>`),
		g.Group(children),
	)
}