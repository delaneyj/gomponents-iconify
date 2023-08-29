package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSwap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.207 15.061L1 11.854v-.707L4.207 7.94l.707.707l-2.353 2.354H15v1H2.56l2.354 2.353l-.707.707zm7.586-7L15 4.854v-.707L11.793.94l-.707.707L13.439 4H1v1h12.44l-2.354 2.354l.707.707z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}