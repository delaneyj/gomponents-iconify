package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5.25 5.125c-.219.01-.433.071-.625.188c-.386.232-.625.65-.625 1.093v13.188c0 .443.239.86.625 1.093c.208.127.449.188.688.188c.2 0 .408-.033.593-.125L12 15.906v3.688c0 .443.239.86.625 1.093c.208.127.448.188.688.188c.2 0 .408-.033.593-.125l8.313-6.594c.4-.376.718-.671.718-1.156c0-.485-.27-.76-.718-1.156L13.906 5.25a1.339 1.339 0 0 0-1.281.063c-.386.232-.625.65-.625 1.093v3.688L5.906 5.25a1.35 1.35 0 0 0-.656-.125z"/>`),
		g.Group(children),
	)
}