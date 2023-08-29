package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rightspeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#ADB8BC" d="m192.585 62.546l151.318 153.699c18.786 19.082 18.786 49.708 0 68.79l-151.318 153.7c-30.772 31.257-83.962 9.467-83.962-34.395V96.941c-.001-43.862 53.189-65.651 83.962-34.395z"/><path fill="#597B91" d="M271.647 142.238h89.869c15.84 0 28.8 12.96 28.8 28.8v159.206c0 15.84-12.96 28.8-28.8 28.8h-89.869V142.238z"/>`),
		g.Group(children),
	)
}