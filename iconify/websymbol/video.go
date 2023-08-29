package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1360 1v1000H0V1h1360zm-200 160h120V81h-120v80zm-200 0h120V81H960v80zm-240 0h120V81H720v80zm240 600h320V241H960v520zM520 161h120V81H520v80zm640 760h120v-80h-120v80zM280 161h120V81H280v80zm240 600h320V241H520v520zm440 160h120v-80H960v80zM80 161h120V81H80v80zm640 760h120v-80H720v80zM80 761h320V241H80v520zm440 160h120v-80H520v80zm-240 0h120v-80H280v80zm-200 0h120v-80H80v80z"/>`),
		g.Group(children),
	)
}