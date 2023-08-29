package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fff" d="M92.9 27.87H33.54c-.57 0-1.04.46-1.04 1.04v12.94c0 .57.47 1.04 1.04 1.04h34.07l-22.84 59.94c-.12.32-.08.68.11.96c.19.28.51.45.86.45H64.6c.43 0 .82-.26.98-.67l28.3-74.3c.12-.32.08-.68-.11-.96c-.21-.27-.53-.44-.87-.44z"/>`),
		g.Group(children),
	)
}