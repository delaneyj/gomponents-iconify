package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaxingGibbousMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WaxingGibbousMoon0" d="M127.76 64c0 35.22-28.54 63.76-63.77 63.76C28.79 127.76.24 99.22.24 64S28.79.24 64 .24c35.22 0 63.76 28.54 63.76 63.76z"/></defs><use fill="#2f2f2f" href="#notoV1WaxingGibbousMoon0"/><clipPath id="notoV1WaxingGibbousMoon1"><use href="#notoV1WaxingGibbousMoon0"/></clipPath><path fill="#fcc21b" d="M142.84 64c0 35.22-25.83 63.76-57.68 63.76c-31.86 0-57.69-28.54-57.69-63.76S53.29.23 85.16.23c31.85.01 57.68 28.55 57.68 63.77z" clip-path="url(#notoV1WaxingGibbousMoon1)"/>`),
		g.Group(children),
	)
}