package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionLtrAcTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M12.5 4.75a.75.75 0 0 0-1.5 0v8.5a.75.75 0 0 0 1.5 0V9.5h1.25a.75.75 0 0 0 0-1.5H12.5V4.75zM3.75 5a.75.75 0 1 0 0 1.5h4.245c-.162 1.634-1.328 4.092-4.46 5.032a.75.75 0 0 0 .43 1.436c4.39-1.317 5.726-5.15 5.532-7.286A.75.75 0 0 0 8.75 5h-5zm13.97 1.22a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 1 1-1.06-1.06l.72-.72h-2.19a.75.75 0 0 1 0-1.5h2.19l-.72-.72a.75.75 0 0 1 0-1.06zm0 8a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 1 1-1.06-1.06l.72-.72H3.75a.75.75 0 0 1 0-1.5h14.69l-.72-.72a.75.75 0 0 1 0-1.06z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}