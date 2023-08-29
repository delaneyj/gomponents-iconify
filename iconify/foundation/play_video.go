package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M85.527 80.647a4.971 4.971 0 0 0 4.973-4.974V24.327a4.971 4.971 0 0 0-4.973-4.974H14.474A4.972 4.972 0 0 0 9.5 24.327v51.346a4.972 4.972 0 0 0 4.974 4.974h71.053zm-4.974-9.948H19.446V29.301h61.107v41.398z"/><path fill="currentColor" d="m64.819 50.288l-11.98 6.913l-11.974 6.917V36.462l11.974 6.918z"/>`),
		g.Group(children),
	)
}