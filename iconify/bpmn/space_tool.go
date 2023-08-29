package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M731.39 100v1800h76V100zm466.824 0v1800h76V100zm401.938 574v216.91l-189.083 1.943v211.667h189.084V1326l325.916-326l-325.916-326zm-1190.834 0L81.29 1000l328.03 326v-221.951h186.97V895.008H409.318V674z"/>`),
		g.Group(children),
	)
}