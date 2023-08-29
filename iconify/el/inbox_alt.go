package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm-73.242 216.064h146.484v168.237h85.474L600 622.412l-158.716-238.11h85.474V216.064zM239.502 589.6h236.865c0 68.266 55.367 123.706 123.633 123.706s123.633-55.44 123.633-123.706h236.865v347.461H239.502V589.6z"/>`),
		g.Group(children),
	)
}