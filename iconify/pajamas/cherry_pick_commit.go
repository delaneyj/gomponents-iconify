package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CherryPickCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.75.75a.75.75 0 0 0-1.5 0v2.528a10.555 10.555 0 0 0-2.55.51a9.11 9.11 0 0 0-.952.375a6.415 6.415 0 0 0-.337.167l-.022.012l-.007.004l-.002.002h-.001L3.75 5l-.372-.651A.75.75 0 0 0 3 5v2.25a.75.75 0 0 0 1.5 0V5.47A9.002 9.002 0 0 1 8 4.75a9.001 9.001 0 0 1 3.5.72v1.78a.75.75 0 0 0 1.5 0V5a.75.75 0 0 0-.378-.651L12.25 5l.372-.651l-.002-.001l-.002-.002l-.007-.004l-.022-.012a6.42 6.42 0 0 0-.337-.168a9.315 9.315 0 0 0-.952-.374a10.555 10.555 0 0 0-2.55-.51V.75ZM10.5 12a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm1.43.75a4.001 4.001 0 0 1-7.86 0H.75a.75.75 0 0 1 0-1.5h3.32a4.001 4.001 0 0 1 7.86 0h3.32a.75.75 0 0 1 0 1.5h-3.32Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}