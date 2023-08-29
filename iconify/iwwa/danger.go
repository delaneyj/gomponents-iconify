package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Danger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M37.874 37.2H2.132a1.742 1.742 0 0 1-1.501-.866a1.707 1.707 0 0 1-.001-1.729L18.507 3.663a1.728 1.728 0 0 1 2.991-.004l17.833 30.869a1.733 1.733 0 0 1-1.457 2.672zM20.001 3.8a.738.738 0 0 0-.628.363L1.496 35.104a.699.699 0 0 0-.001.727a.739.739 0 0 0 .637.369h35.742a.741.741 0 0 0 .732-.73a.787.787 0 0 0-.134-.429L20.634 4.163a.738.738 0 0 0-.633-.363z"/><path fill="currentColor" d="M20.002 27.948a.5.5 0 0 1-.5-.5V9.864a.5.5 0 0 1 1 0v17.584a.5.5 0 0 1-.5.5zm-.004 3.688a.5.5 0 0 1-.5-.5V29.78a.5.5 0 0 1 1 0v1.355a.5.5 0 0 1-.5.501z"/>`),
		g.Group(children),
	)
}