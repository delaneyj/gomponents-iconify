package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m8.645 23.722l22.062 5.288c.215.061.439.069.67.023c.232-.046.424-.146.578-.3l7.4-6.16a4.647 4.647 0 0 0-1.664-1.701a4.469 4.469 0 0 0-2.36-.644h-22.2c-1.079 0-2.034.33-2.867.988a4.47 4.47 0 0 0-1.619 2.506h0Z"/><path d="M8.645 23.722V17.28a4.868 4.868 0 0 1 4.868-4.869h20.974a4.868 4.868 0 0 1 4.869 4.869v5.293"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.961 9.286H35.04a7.465 7.465 0 0 1 7.461 7.462v14.566c0 4.084-3.316 7.4-7.4 7.4H12.9a7.404 7.404 0 0 1-7.4-7.4V16.747a7.465 7.465 0 0 1 7.461-7.462Z"/>`),
		g.Group(children),
	)
}