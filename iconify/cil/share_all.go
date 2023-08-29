package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M242.918 16.173h-42.361v88H152A136.268 136.268 0 0 0 28.025 184.2A114.159 114.159 0 0 0 16 235.506v260.667h52.146l49.606-177.756a85.4 85.4 0 0 1 81.993-62.244h.812v88h41.78l164.29-164.29ZM232.557 308.7v-84.527h-32.812A117.5 117.5 0 0 0 86.93 309.815L48 449.315V235.506a82.454 82.454 0 0 1 8.785-37.276l.292-.614A104.217 104.217 0 0 1 152 136.173h80.557V51.067l128.816 128.816Z"/><path fill="currentColor" d="M330.918 15.509h-43.409l164.373 164.374l-163.626 163.626h42.081l164.29-164.29l-163.709-163.71z"/>`),
		g.Group(children),
	)
}