package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prestashop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 960h-64V800q0-32-19-57t-49-34q0-30-31.5-49.5T608 640h-32v-32q0-40-28-68t-68-28t-68 28t-28 68v102q-28 10-46 34.5T320 800v160H64q-26 0-45-18.5T0 896V320q0-26 18.5-45T64 256h64V128L256 0v256h384V0l128 128v128h64q27 0 45.5 19t18.5 45v576q0 27-18.5 45.5T832 960zM320 0h256v128H320V0zm96 768q13 0 22.5 9.5T448 800V608q0-13 9.5-22.5T480 576t22.5 9.5T512 608v128q0-13 9.5-22.5T544 704t22.5 9.5T576 736q0-13 9.5-22.5T608 704t22.5 9.5T640 736v64q0-13 9.5-22.5T672 768t22.5 9.5T704 800v96q0 53-31 90.5t-76 37.5H491q-44 0-75.5-37.5T384 896v-96q0-13 9.5-22.5T416 768z"/>`),
		g.Group(children),
	)
}