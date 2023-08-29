package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiveStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.596 1.164a.75.75 0 0 1 1.054.118a10.75 10.75 0 0 1-.017 13.438a.75.75 0 1 1-1.17-.94a9.25 9.25 0 0 0 .015-11.562a.75.75 0 0 1 .118-1.054Zm-3.84 2.97a.75.75 0 0 1 1.048.166a6.281 6.281 0 0 1 0 7.381a.75.75 0 1 1-1.214-.881a4.781 4.781 0 0 0 0-5.619a.75.75 0 0 1 .167-1.047ZM3 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}