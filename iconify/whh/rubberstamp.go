package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubberstamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 833H64q-27 0-45.5-19T0 769v-64q0-27 18.5-45.5T64 641h346q38-49 38-128q0-73-28-133.5T345 286q-25-44-25-93q0-80 56-136.5T512 0t136 56.5T704 193q0 49-25 93q-47 33-75 93.5T576 513q0 79 38 128h346q27 0 45.5 18.5T1024 705v64q0 26-18.5 45T960 833zm0 128q0 26-18.5 45t-45.5 19H128q-27 0-45.5-19T64 961v-64h896v64z"/>`),
		g.Group(children),
	)
}