package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaningGibbousMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WaningGibbousMoon0" d="M.23 64c0 35.22 28.54 63.77 63.77 63.77c35.21 0 63.77-28.54 63.77-63.77C127.77 28.79 99.22.23 64 .23S.23 28.79.23 64z"/></defs><use fill="#2f2f2f" href="#notoV1WaningGibbousMoon0"/><clipPath id="notoV1WaningGibbousMoon1"><use href="#notoV1WaningGibbousMoon0"/></clipPath><path fill="#fcc21b" d="M-14.84 64c0 35.22 25.82 63.77 57.69 63.77c31.85 0 57.68-28.54 57.68-63.77C100.53 28.79 74.7.23 42.85.23C10.98.23-14.84 28.79-14.84 64z" clip-path="url(#notoV1WaningGibbousMoon1)"/>`),
		g.Group(children),
	)
}