package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M891 451q5 31 5 61t-5 61l129 124q5 5 3.5 15.5T1015 734l-69 107q-7 11-17 17.5t-16 5.5l-165-49q-46 35-98 55l-39 138q-2 6-13 11t-25 5H436q-14 0-24.5-5t-12.5-11l-36-142q-44-19-87-52l-164 50q-7 1-17-5.5T78 841L9 734q-7-11-8.5-21.5T4 697l129-124q-5-31-5-61t5-61L4 327q-5-5-3.5-15.5T9 290l69-107q7-11 17-17.5t17-5.5l164 50q45-36 98-56l39-138q2-6 13-11t25-5h137q14 0 25 5t12 11l36 142q45 19 87 52l165-50q7-1 16.5 5.5T946 183l69 107q7 11 8.5 21.5T1020 327zM512 320q-80 0-136 56.5t-56 136T376 648t136 56t136-56t56-135.5t-56-136T512 320z"/>`),
		g.Group(children),
	)
}