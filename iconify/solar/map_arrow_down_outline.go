package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.047 1.987c.654.685.94 1.768.473 2.816l-7.363 16.51a2.338 2.338 0 0 1-4.315 0L2.48 4.802a2.546 2.546 0 0 1 .473-2.816c.659-.69 1.735-1.009 2.767-.458l-.353.662l.353-.662l5.904 3.152l-.354.662l.354-.662a.789.789 0 0 0 .752 0l5.904-3.15l.353.662l-.353-.662c1.032-.55 2.108-.232 2.767.458Zm-2.06.865l-.351-.656l.35.656l-5.904 3.153a2.289 2.289 0 0 1-2.165 0L5.014 2.852c-.356-.19-.715-.103-.976.171c-.268.28-.388.72-.188 1.169l7.362 16.51c.326.73 1.25.73 1.575 0l7.363-16.51c.2-.448.08-.889-.188-1.169c-.262-.274-.62-.36-.976-.17Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}