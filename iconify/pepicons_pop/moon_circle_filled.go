package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMoonCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M11.764 9.128a4.506 4.506 0 0 1 2.6-3.196c.93-.41.727-1.784-.281-1.908c-4.612-.566-8.921 2.498-9.895 7.078c-1.035 4.87 2.067 9.656 6.93 10.69c4.862 1.033 9.642-2.078 10.677-6.947a9.045 9.045 0 0 0 .195-2.166c-.032-1.019-1.388-1.343-1.877-.449a4.504 4.504 0 0 1-4.885 2.245a4.505 4.505 0 0 1-3.464-5.347Zm-.23 10.707c-3.782-.803-6.195-4.527-5.39-8.317a7.022 7.022 0 0 1 4.727-5.23a6.512 6.512 0 0 0-1.063 2.424c-.747 3.516 1.493 6.973 5.005 7.72a6.487 6.487 0 0 0 4.636-.751c-1.273 3.053-4.57 4.866-7.916 4.155Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMoonCircleFilled0)"/></g>`),
		g.Group(children),
	)
}