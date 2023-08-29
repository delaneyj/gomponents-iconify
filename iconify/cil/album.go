package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294ZM256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48s208 93.309 208 208s-93.309 208-208 208Z"/><path fill="currentColor" d="M256 152a104 104 0 1 0 104 104a104.118 104.118 0 0 0-104-104Zm0 176a72 72 0 1 1 72-72a72.081 72.081 0 0 1-72 72Z"/><path fill="currentColor" d="M240 240h32v32h-32zm16-128V80a174.144 174.144 0 0 0-79.968 19.178A177.573 177.573 0 0 0 115.2 150.39l25.586 19.219A142.923 142.923 0 0 1 256 112Z"/>`),
		g.Group(children),
	)
}