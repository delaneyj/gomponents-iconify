package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dotclear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 1024q-119 0-203.5-84.5T448 736t84.5-203.5T736 448q87 0 160 49v82q-66-67-160-67q-93 0-158.5 65.5T512 736t65.5 158.5T736 960t158.5-65.5T960 736V192h64v544q0 119-84.5 203.5T736 1024zM419 391q-44 29-81 43t-82 14q-54 0-96-32q-10 11-79.5 44.5T0 512v-64q37-43 160-96q71-30 110-47t96-43t87.5-42t63-34.5t46-32.5t13.5-25q-64 32-160 64q-39 13-80.5 31T227 273.5T128 320q0-43 10-78t22-50q27-35 74-56t100.5-30.5t108-19T551 56t89-56q0 96-23.5 165T547 286T419 391z"/>`),
		g.Group(children),
	)
}