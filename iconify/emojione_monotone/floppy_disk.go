package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M36 6h10v16H36z"/><path fill="currentColor" d="m61.293 9.293l-6.586-6.586C54.318 2.318 53.55 2 53 2H9v4H5V2H3c-.55 0-1 .45-1 1v58c0 .55.45 1 1 1h58c.55 0 1-.45 1-1V11c0-.55-.318-1.318-.707-1.707M52 58H12V34c0-1.1.9-2 2-2h36c1.1 0 2 .9 2 2v24m0-36c0 1.1-.899 2-2 2H20c-1.1 0-2-.9-2-2V3h34v19m7 36c0 .55-.45 1-1 1h-2c-.55 0-1-.45-1-1v-2c0-.55.45-1 1-1h2c.55 0 1 .45 1 1v2"/><path fill="currentColor" d="M17 36h30v2H17zm0 6h30v2H17zm0 6h30v2H17z"/>`),
		g.Group(children),
	)
}