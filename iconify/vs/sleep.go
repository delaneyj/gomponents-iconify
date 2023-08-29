package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1152 992v64q0 14-9 23t-23 9H672q-14 0-23-9t-9-23v-64q0-32 20-56l300-360H672q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56L832 960h288q14 0 23 9t9 23zm640 384v64q0 14-9 23t-23 9h-448q-14 0-23-9t-9-23v-64q0-32 20-56l300-360h-288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56l-300 360h288q14 0 23 9t9 23zM512 608v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-32 20-56l300-360H32q-14 0-23-9t-9-23V96q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56L192 576h288q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}