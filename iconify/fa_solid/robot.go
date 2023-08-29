package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Robot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 224h32v192H32a31.962 31.962 0 0 1-32-32V256a31.962 31.962 0 0 1 32-32Zm512-48v272a64.063 64.063 0 0 1-64 64H160a64.063 64.063 0 0 1-64-64V176a79.974 79.974 0 0 1 80-80h112V32a32 32 0 0 1 64 0v64h112a79.974 79.974 0 0 1 80 80Zm-280 80a40 40 0 1 0-40 40a39.997 39.997 0 0 0 40-40Zm-8 128h-64v32h64Zm96 0h-64v32h64Zm104-128a40 40 0 1 0-40 40a39.997 39.997 0 0 0 40-40Zm-8 128h-64v32h64Zm192-128v128a31.962 31.962 0 0 1-32 32h-32V224h32a31.962 31.962 0 0 1 32 32Z"/>`),
		g.Group(children),
	)
}