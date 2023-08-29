package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinnTheHumanBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 88H96a52 52 0 0 0 0 104h64a52 52 0 0 0 0-104Zm0 80H96a28 28 0 0 1 0-56h64a28 28 0 0 1 0 56Zm-48-28a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm64 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm36-112a36 36 0 0 0-33.94 24H77.94A36 36 0 0 0 8 64v76a88.1 88.1 0 0 0 88 88h64a88.1 88.1 0 0 0 88-88V64a36 36 0 0 0-36-36Zm12 112a64.07 64.07 0 0 1-64 64H96a64.07 64.07 0 0 1-64-64V64a12 12 0 0 1 24 0a12 12 0 0 0 12 12h120a12 12 0 0 0 12-12a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}