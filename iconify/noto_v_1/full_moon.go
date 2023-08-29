package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1FullMoon0" d="M.23 64c0 35.21 28.54 63.76 63.77 63.76c35.21 0 63.77-28.55 63.77-63.76C127.77 28.78 99.22.23 64 .23S.23 28.78.23 64z"/></defs><use fill="#fcc21b" href="#notoV1FullMoon0"/>`),
		g.Group(children),
	)
}