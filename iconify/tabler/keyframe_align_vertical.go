package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyframeAlignVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v2m.816 12.58c-.207.267-.504.42-.816.42c-.312 0-.61-.153-.816-.42l-2.908-3.748a1.39 1.39 0 0 1 0-1.664l2.908-3.748c.207-.267.504-.42.816-.42c.312 0 .61.153.816.42l2.908 3.748a1.39 1.39 0 0 1 0 1.664l-2.908 3.748zM12 20v2"/>`),
		g.Group(children),
	)
}