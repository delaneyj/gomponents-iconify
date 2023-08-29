package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.06 6l-.21-.2c-.52-.54-.43-.79.08-1.3l2.72-2.75c.81-.82.96-1.21 1.73-.48l.21.2zm.53.45l4.4-4.4c.7.94 2.34 3.47 1.53 5.34c-.73 1.67-1.09 1.75-2 3c-1.85 2.11-4.18 4.37-6 6.07c-1.26.91-1.31 1.33-3 2c-1.8.71-4.4-.89-5.38-1.56l4.4-4.4l1.18 1.62c.34.46 1.2-.06 1.8-.66c1.04-1.05 3.18-3.18 4-4.07c.59-.59 1.12-1.45.66-1.8zM1.57 16.5l-.21-.21c-.68-.74-.29-.9.52-1.7l2.74-2.72c.51-.49.75-.6 1.27-.11l.2.21z"/>`),
		g.Group(children),
	)
}