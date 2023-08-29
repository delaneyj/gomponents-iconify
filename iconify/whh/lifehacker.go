package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifehacker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M672 1024h-64q-13 0-22.5-9.5T576 992V480q0-45-36.5-70.5T448 384q-41 0-64 14v594q0 13-9.5 22.5T352 1024h-64q-13 0-22.5-9.5T256 992V32q0-13 9.5-22.5T288 0h64q13 0 22.5 9.5T384 32v237q24-13 64-13q107 0 181.5 64.5T704 480v512q0 13-9.5 22.5T672 1024zm-512 0H32q-13 0-22.5-9.5T0 992V32Q0 19 9.5 9.5T32 0h128q13 0 22.5 9.5T192 32v960q0 13-9.5 22.5T160 1024z"/>`),
		g.Group(children),
	)
}