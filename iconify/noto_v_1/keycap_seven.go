package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fcc21b" d="M116 128H12c-6.6 0-12-5.4-12-12V12C0 5.4 5.4 0 12 0h104c6.6 0 12 5.4 12 12v104c0 6.6-5.4 12-12 12z"/><path fill="#f79329" d="M124.48 3.52L3.52 124.48C5.7 126.65 8.7 128 12 128h104c6.6 0 12-5.4 12-12V12c0-3.3-1.35-6.3-3.52-8.48z"/><path fill="#fcc21b" d="M112 101c0 6.63-5.37 12-12 12H27.89c-6.6 0-12-5.4-12-12V28.9c0-6.6 5.4-12 12-12H100c6.63 0 12 5.37 12 12V101z"/><path fill="#fff" d="M92.9 27.87H33.54c-.57 0-1.04.46-1.04 1.04v12.94c0 .57.47 1.04 1.04 1.04h34.07l-22.84 59.94c-.12.32-.08.68.11.96c.19.28.51.45.86.45H64.6c.43 0 .82-.26.98-.67l28.3-74.3c.12-.32.08-.68-.11-.96c-.21-.27-.53-.44-.87-.44z"/>`),
		g.Group(children),
	)
}