package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThunderstormScattered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m13.338 30l-1.736-1l2.287-4H10l3.993-7l1.737 1l-2.284 4h3.891l-3.999 7z"/><path fill="currentColor" d="M27 11a5.995 5.995 0 0 0-11.367-2.669A7.35 7.35 0 0 0 13.5 8a7.55 7.55 0 0 0-7.15 5.244A5.993 5.993 0 0 0 8 25v-2a3.993 3.993 0 0 1-.673-7.93l.663-.112l.146-.656a5.496 5.496 0 0 1 10.73 0l.145.656l.663.113A3.993 3.993 0 0 1 19 23v2a5.985 5.985 0 0 0 5.216-8.946A5.994 5.994 0 0 0 27 11Zm-4.073 3.49a5.96 5.96 0 0 0-2.278-1.246a7.516 7.516 0 0 0-3.185-4.069a3.987 3.987 0 1 1 5.463 5.314ZM28.586 19l-2.122-2.12l1.414-1.415L30 17.587zM29 10h3v2h-3zm-2.536-4.878L28.586 3L30 4.414l-2.122 2.122zM20 0h2v3h-2zm-5.878 6.536L12 4.414L13.414 3l2.122 2.122z"/>`),
		g.Group(children),
	)
}