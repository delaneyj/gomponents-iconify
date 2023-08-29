package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Algorhythm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1015 565L896 684v244q0 13-9.5 22.5T864 960H481q-30 29-81 46.5T288 1024q-93 0-158.5-37.5T64 896q0-48 55-83.5T256 770V640H64q-26 0-45-19T0 576V448q0-27 19-45.5T64 384h192V254q-82-7-137-42.5T64 128q0-53 65.5-90.5T288 0t158.5 37.5T512 128q0 48-55 83.5T320 254v130h192q27 0 45.5 18.5T576 448v64h146l120-120q9-9 21.5-9t21.5 9l130 130q9 9 9 21.5t-9 21.5zm-503 75H320v130q82 7 137 42.5t55 83.5h320V685L723 576H576q0 26-18.5 45T512 640z"/>`),
		g.Group(children),
	)
}