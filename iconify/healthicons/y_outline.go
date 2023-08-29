package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M31.695 10.08a3 3 0 0 1 4.61 3.84l-.769-.64l.769.64L27 25.087V36a3 3 0 0 1-6 0V25.086l-9.305-11.165a3 3 0 0 1 .384-4.226l.64.769l-.64-.769a3 3 0 0 1 4.226.384L24 19.314l7.695-9.235Zm2.945 1.152a1 1 0 0 0-1.408.128l-8.464 10.156a1 1 0 0 1-1.536 0L14.768 11.36a1 1 0 1 0-1.536 1.28l9.536 11.444a1 1 0 0 1 .232.64V36a1 1 0 1 0 2 0V24.724a1 1 0 0 1 .232-.64l9.536-11.444a1 1 0 0 0-.128-1.408Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}