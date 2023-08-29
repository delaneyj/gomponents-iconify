package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miraheze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.677 12.923l3.768.011l1.949 3.369l-1.926 3.323H2.666L.727 16.292l1.95-3.369Zm-.004-8.6l3.761.011l1.944 3.367l-1.922 3.326H2.662L.727 7.69l1.946-3.367Zm14.882 0l3.768.011l1.95 3.367l-1.928 3.326h-3.801L15.606 7.69l1.949-3.367Zm0 8.6l3.768.011l1.95 3.369l-1.928 3.323h-3.802l-1.937-3.334l1.949-3.369Zm-7.456 4.373l3.767.011l1.951 3.368L13.889 24h-3.801l-1.939-3.336l1.95-3.368Zm0-17.296l3.767.011l1.951 3.369l-1.928 3.324h-3.801L8.149 3.368L10.099 0Zm0 8.628l3.767.011l1.951 3.368l-1.928 3.325h-3.801l-1.939-3.336l1.95-3.368Z"/>`),
		g.Group(children),
	)
}