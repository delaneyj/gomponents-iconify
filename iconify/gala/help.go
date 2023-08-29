package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaHelp0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><circle id="galaHelp1" cx="181.018" cy="-.001" r="112.049" stroke-opacity="1" transform="rotate(45)"/><circle id="galaHelp2" cx="181.018" cy="-.001" r="48.021" transform="rotate(45)"/><path id="galaHelp3" stroke-opacity="1" d="M 171.33275,107.30371 226.85481,75.248027"/><path id="galaHelp4" stroke-opacity="1" d="M 148.69466,84.66561 180.75035,29.143566"/><path id="galaHelp5" stroke-opacity="1" d="M 84.667256,148.69302 29.145241,180.74869"/><path id="galaHelp6" stroke-opacity="1" d="M 107.30535,171.33111 75.249699,226.85316"/><path id="galaHelp7" stroke-opacity="1" d="m 148.69466,171.33111 32.05569,55.52205"/><path id="galaHelp8" stroke-opacity="1" d="m 171.33275,148.69302 55.52206,32.05567"/><path id="galaHelp9" stroke-opacity="1" d="M 107.30535,84.665618 75.249699,29.143573"/><path id="galaHelpa" stroke-opacity="1" d="M 84.667256,107.30372 29.145241,75.248046"/></g>`),
		g.Group(children),
	)
}