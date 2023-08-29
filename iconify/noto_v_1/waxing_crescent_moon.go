package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaxingCrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WaxingCrescentMoon0" d="M.24 64.01c0 35.21 28.54 63.76 63.76 63.76c35.21 0 63.77-28.55 63.77-63.76C127.77 28.79 99.22.23 64 .23S.24 28.78.24 64.01z"/></defs><use fill="#fcc21b" href="#notoV1WaxingCrescentMoon0"/><clipPath id="notoV1WaxingCrescentMoon1"><use href="#notoV1WaxingCrescentMoon0"/></clipPath><path fill="#2f2f2f" d="M-14.84 64.01c0 35.21 25.82 63.76 57.69 63.76c31.85 0 57.68-28.55 57.68-63.76C100.53 28.79 74.7.23 42.85.23c-31.87 0-57.69 28.55-57.69 63.78z" clip-path="url(#notoV1WaxingCrescentMoon1)"/>`),
		g.Group(children),
	)
}