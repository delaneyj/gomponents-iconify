package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.19 7.45A3.49 3.49 0 0 0 13 4.32C13 1.65 10.93 0 7.91 0H2.06v16h6.19c3.5 0 5.69-1.76 5.69-4.69a4 4 0 0 0-2.75-3.86zm-7.88-6.2h4.6a4.36 4.36 0 0 1 2.86.87a2.71 2.71 0 0 1 1 2.2c0 .94-.43 2.51-3.36 2.51h-5.1zm8.31 12.54a5.3 5.3 0 0 1-3.37 1H3.31V8.08h4.35a7.29 7.29 0 0 1 3.68.84a2.62 2.62 0 0 1 1.35 2.39a3.05 3.05 0 0 1-1.07 2.48z"/>`),
		g.Group(children),
	)
}