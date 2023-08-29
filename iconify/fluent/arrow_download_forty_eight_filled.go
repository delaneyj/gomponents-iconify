package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownloadFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.25 38.5h23.5a1.75 1.75 0 0 1 .144 3.494L35.75 42h-23.5a1.75 1.75 0 0 1-.143-3.494l.143-.006h23.5h-23.5ZM23.607 6.256l.143-.006a1.75 1.75 0 0 1 1.744 1.606L25.5 8v21.333l4.793-4.792a1.75 1.75 0 0 1 2.475 2.475l-7.778 7.778a1.75 1.75 0 0 1-2.475 0l-7.778-7.778a1.75 1.75 0 0 1 2.475-2.475L22 29.329V8a1.75 1.75 0 0 1 1.607-1.744l.143-.006l-.143.006Z"/>`),
		g.Group(children),
	)
}