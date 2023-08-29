package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1FirstQuarterMoon0" d="M127.76 64c0 35.22-28.54 63.77-63.77 63.77C28.79 127.76.24 99.22.24 64C.24 28.79 28.79.24 64 .24c35.22 0 63.76 28.55 63.76 63.76z"/></defs><use fill="#2f2f2f" href="#notoV1FirstQuarterMoon0"/><clipPath id="notoV1FirstQuarterMoon1"><use href="#notoV1FirstQuarterMoon0"/></clipPath><path fill="#fcc21b" d="M64-17.32h88.89v158.99H64z" clip-path="url(#notoV1FirstQuarterMoon1)"/>`),
		g.Group(children),
	)
}