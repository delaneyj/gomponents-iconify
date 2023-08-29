package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1LastQuarterMoon0" d="M.24 64c0 35.21 28.54 63.77 63.76 63.77c35.21 0 63.77-28.55 63.77-63.77S99.22.23 64 .23S.24 28.78.24 64z"/></defs><use fill="#2f2f2f" href="#notoV1LastQuarterMoon0"/><clipPath id="notoV1LastQuarterMoon1"><use href="#notoV1LastQuarterMoon0"/></clipPath><path fill="#fcc21b" d="M-24.89-17.32H64v158.99h-88.89z" clip-path="url(#notoV1LastQuarterMoon1)"/>`),
		g.Group(children),
	)
}