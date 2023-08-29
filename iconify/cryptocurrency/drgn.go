package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drgn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.403 9.957L9.4 20.78h2.818l-.072-7.327l8.28 10.513l2.228 2.83l-.018-1.573l-.029-14.179h-2.774l.072 7.386L9.577 5.306l-.173-.22v3.948l-.004.92l.003.003zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16z"/><path fill="currentColor" d="m9.403 9.957l11.024 14.01l-8.281-10.514l.072 7.327H9.4l.003-10.823zm0-.923l.001-3.947l.173.219l-.161-.204l-.013 3.932zm10.502 9.396l-.072-7.386h2.774l.03 14.18l-.038-3.37l-2.694-3.424z" opacity=".398"/>`),
		g.Group(children),
	)
}