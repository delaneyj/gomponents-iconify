package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 0C28.7 0 0 28.7 0 64v288c0 88.4 71.6 160 160 160s160-71.6 160-160V64c0-35.3-28.7-64-64-64H64zm96 416a48 48 0 1 1 0-96a48 48 0 1 1 0 96zm48-176a48 48 0 1 1-96 0a48 48 0 1 1 96 0zm-48-80a48 48 0 1 1 0-96a48 48 0 1 1 0 96z"/>`),
		g.Group(children),
	)
}