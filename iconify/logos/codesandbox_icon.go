package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodesandboxIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M256 0v256H0V0h256Zm-26.182 26.182H26.182v203.636h203.636V26.182Z"/>`),
		g.Group(children),
	)
}