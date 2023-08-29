package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 136H128v152h80v88h176V224h-80ZM160 256v-88h112v88H160Zm192 0v88H240v-56h64v-32Z"/><path fill="currentColor" d="M400 48H112V16H16v96h32v288H16v96h96v-32h288v32h96v-96h-32V112h32V16h-96ZM48 48h32v32H48Zm32 416H48v-32h32Zm384 0h-32v-32h32ZM432 48h32v32h-32Zm0 352h-32v32H112v-32H80V112h32V80h288v32h32Z"/>`),
		g.Group(children),
	)
}