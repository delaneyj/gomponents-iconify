package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 134.702h51.287l90.31 154.728l90.312-152.465h51.287L141.598 377.298L0 134.702zm218.147 234.382l23.235-40.05h102.495l-25.315 42.313h51.472L416.32 294.8l45.787 77.68H512l-70.286-120.757l67.444-111.284h-52.103l-40.419 64.997l-38.84-67.26h-50.84l63.471 111.152l-10.736 19.262l-74.523-130.73l-137.545 231.225h50.524zm51.385-83.72l34.573-59.333l33.638 59.333h-68.21z"/>`),
		g.Group(children),
	)
}