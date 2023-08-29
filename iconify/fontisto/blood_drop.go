package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.647 13.663a1.923 1.923 0 0 0-.104-.239l.005.011c-.146-.341-1.32-2.609-1.516-2.918L7.581-.001l-5.59 10.776c-.19.31-1.232 2.32-1.376 2.66l-.057.126A7.456 7.456 0 0 0 0 16.417a7.582 7.582 0 1 0 15.164 0v-.006l.001-.101c0-.955-.19-1.866-.535-2.696l.017.047zm-7.062 8.355a5.585 5.585 0 0 1-4.899-8.269l-.015.029a6.382 6.382 0 0 0-.11 1.181v.001a6.7 6.7 0 0 0 6.696 6.696c.113 0 .234 0 .346-.006a5.53 5.53 0 0 1-2.003.368h-.023h.001z"/>`),
		g.Group(children),
	)
}