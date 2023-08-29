package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Red(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 142.48l71.37-36.977L256 0l-71.37 105.503L256 142.481zM256 512l256-133.172l-71.37-105.503L256 369.52L71.37 273.325L0 378.828L256 512zm0-184.63l163.685-85.334l-71.37-105.503L256 184.63l-92.315-48.097l-71.37 105.503L256 327.37z"/>`),
		g.Group(children),
	)
}