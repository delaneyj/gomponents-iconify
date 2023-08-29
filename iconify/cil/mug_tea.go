package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MugTea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M436.574 120H352V64H32v344a64.072 64.072 0 0 0 64 64h192a64.072 64.072 0 0 0 64-64v-8h84.574A59.493 59.493 0 0 0 496 340.574V179.426A59.493 59.493 0 0 0 436.574 120Zm-275.2 118.894L192 269.521v57.373h-64v-57.373l30.627-30.627ZM464 340.574A27.457 27.457 0 0 1 436.574 368H320v40a32.036 32.036 0 0 1-32 32H96a32.036 32.036 0 0 1-32-32V96h80v112.266l-48 48v102.628h128V256.266l-48-48V96h144v56h116.574A27.457 27.457 0 0 1 464 179.426Z"/><path fill="currentColor" d="M416 180h-96v160h96a20.023 20.023 0 0 0 20-20V200a20.023 20.023 0 0 0-20-20Zm-12 128h-52v-96h52Z"/>`),
		g.Group(children),
	)
}