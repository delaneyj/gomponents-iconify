package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minitonguealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M951.443 732q0 122-75 207.5t-181 85.5t-181-85.5t-75-207.5h-366q-30 0-51.5-21t-21.5-51.5t21.5-52t51.5-21.5h878q30 0 51.5 21.5t21.5 52t-21.5 51.5t-51.5 21zm-366 0q0 61 32 104t77.5 43t78-43t32.5-104h-220zm350.5-383q-15.5 17-38 17t-37.5-17l-167-111q-12-4-20-13q-15-17-15-41.5t15-41.5q8-9 20-14l167-111q15-17 37.5-17t38 17.5t15.5 41.5t-16 41l-124 83l124 83q16 18 16 42t-15.5 41zm-604.5-111l-167 111q-15 17-37.5 17t-38-17t-15.5-41t16-42l124-83l-124-83q-16-17-16-41t15.5-41.5t38-17.5t37.5 17l167 111q12 5 20 14q15 17 15 41.5t-15 41.5q-8 9-20 13z"/>`),
		g.Group(children),
	)
}