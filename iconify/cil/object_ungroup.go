package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectUngroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 168h-72v-56h32V16h-96v32H112V16H16v96h32v152H16v96h96v-32h72v72h-32v96h96v-32h152v32h96v-96h-32V232h32v-96h-96ZM296 48h32v32h-32Zm32 248v32h-32v-32ZM48 48h32v32H48Zm32 280H48v-32h32Zm32-32v-32H80V112h32V80h152v32h32v152h-32v32Zm104 168h-32v-32h32Zm248 0h-32v-32h32Zm-32-296h32v32h-32Zm0 232h-32v32H248v-32h-32v-72h48v32h96v-96h-32v-64h72v32h32Z"/>`),
		g.Group(children),
	)
}