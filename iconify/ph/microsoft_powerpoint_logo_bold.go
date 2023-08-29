package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPowerpointLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M96 92H80a12 12 0 0 0-12 12v48a12 12 0 0 0 24 0v-4h4a28 28 0 0 0 0-56Zm0 32h-4v-8h4a4 4 0 0 1 0 8Zm40-104a107.49 107.49 0 0 0-80.5 36H36a20 20 0 0 0-20 20v104a20 20 0 0 0 20 20h19.5A108 108 0 1 0 136 20Zm0 24a84.12 84.12 0 0 1 83.13 72H168V76a20 20 0 0 0-20-20H92.68A83.67 83.67 0 0 1 136 44ZM40 80h104v96H40Zm96 132a83.67 83.67 0 0 1-43.32-12H148a20 20 0 0 0 20-20v-40h51.13A84.12 84.12 0 0 1 136 212Z"/>`),
		g.Group(children),
	)
}