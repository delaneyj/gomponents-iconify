package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm96-768H352q-13 0-22.5 9.5T320 288t9.5 22.5T352 320h96v416q0 13 9.5 22.5T480 768t22.5-9.5T512 736V320h96q13 0 22.5-9.5T640 288t-9.5-22.5T608 256z"/>`),
		g.Group(children),
	)
}