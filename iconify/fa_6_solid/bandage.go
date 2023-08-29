package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bandage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 416h96c35.3 0 64-28.7 64-64V160c0-35.3-28.7-64-64-64h-96v320zM448 96H192v320h256V96zM64 96c-35.3 0-64 28.7-64 64v192c0 35.3 28.7 64 64 64h96V96H64zm184 112a24 24 0 1 1 48 0a24 24 0 1 1-48 0zm120-24a24 24 0 1 1 0 48a24 24 0 1 1 0-48zM248 304a24 24 0 1 1 48 0a24 24 0 1 1-48 0zm120-24a24 24 0 1 1 0 48a24 24 0 1 1 0-48z"/>`),
		g.Group(children),
	)
}