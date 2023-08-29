package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagramSuccessor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 416v-64c0-35.3-28.7-64-64-64H64c-35.3 0-64 28.7-64 64v64c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64zM64 160V96h160v64H64zm224 0V96h80c8.8 0 16 7.2 16 16v16h-38.1c-21.4 0-32.1 25.9-17 41l70.1 70c9.4 9.4 24.6 9.4 33.9 0l70.1-70c15.1-15.1 4.4-41-17-41h-38v-16c0-44.2-35.8-80-80-80H64C28.7 32 0 60.7 0 96v64c0 35.3 28.7 64 64 64h160c35.3 0 64-28.7 64-64z"/>`),
		g.Group(children),
	)
}