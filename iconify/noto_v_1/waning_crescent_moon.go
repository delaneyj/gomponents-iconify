package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaningCrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WaningCrescentMoon0" d="M127.76 64c0 35.22-28.55 63.77-63.76 63.77C28.78 127.77.23 99.22.23 64C.23 28.79 28.78.23 64 .23c35.21 0 63.76 28.56 63.76 63.77z"/></defs><use fill="#fcc21b" href="#notoV1WaningCrescentMoon0"/><clipPath id="notoV1WaningCrescentMoon1"><use href="#notoV1WaningCrescentMoon0"/></clipPath><path fill="#2f2f2f" d="M142.83 64c0 35.22-25.82 63.77-57.68 63.77S27.46 99.23 27.46 64C27.46 28.79 53.28.23 85.15.23c31.86 0 57.68 28.56 57.68 63.77z" clip-path="url(#notoV1WaningCrescentMoon1)"/>`),
		g.Group(children),
	)
}