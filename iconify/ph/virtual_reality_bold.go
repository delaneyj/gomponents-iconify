package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualRealityBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m123.12 100.52l-26 64a12 12 0 0 1-22.24 0l-26-64a12 12 0 0 1 22.24-9L86 128.12l14.88-36.64a12 12 0 0 1 22.24 9ZM256 128a84.09 84.09 0 0 1-84 84H84a84 84 0 0 1 0-168h88a84.09 84.09 0 0 1 84 84Zm-24 0a60.07 60.07 0 0 0-60-60H84a60 60 0 0 0 0 120h88a60.07 60.07 0 0 0 60-60Zm-42 11.24l8.46 14.81A12 12 0 1 1 177.58 166l-10.26-18H160v12a12 12 0 0 1-24 0V96a12 12 0 0 1 12-12h20a32 32 0 0 1 22 55.24ZM160 124h8a8 8 0 0 0 0-16h-8Z"/>`),
		g.Group(children),
	)
}