package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TorresStraitIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTorresStraitIslands0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTorresStraitIslands0)"><path fill="#0052b4" d="m0 128l256-32l256 32v256l-256 32L0 384Z"/><path fill="#333" d="m0 96l256-32l256 32v32H0Z"/><path fill="#6da544" d="M0 0h512v96H0Z"/><path fill="#333" d="m0 416l256 32l256-32v-32H0Z"/><path fill="#6da544" d="M0 512h512v-96H0Z"/><path fill="#eee" d="M245 144c-106 32-101 112-67 186l-40 38l73-27v-98c24-35 66-35 90 0v98l73 27l-40-38c34-74 39-154-67-186l-11 46l-11-46zm11 83l-9 27h-28l23 17l-9 28l23-17l23 17l-9-28l24-17h-29l-9-27z"/></g>`),
		g.Group(children),
	)
}