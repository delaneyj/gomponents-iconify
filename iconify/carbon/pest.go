package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="7.5" cy="9.5" r="1.5" fill="currentColor"/><circle cx="13" cy="13" r="1" fill="currentColor"/><circle cx="22" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M25 14a8.984 8.984 0 0 0-7 3.356V13a10.905 10.905 0 0 0-1.01-4.585l-.356-.776l-.823.231A2.955 2.955 0 0 1 15 8a3.003 3.003 0 0 1-3-3a2.957 2.957 0 0 1 .13-.81l.232-.824l-.777-.357A10.912 10.912 0 0 0 7 2H2v7a10.895 10.895 0 0 0 2.234 6.627l.389.428l.54-.038a5.529 5.529 0 0 1 5.362 3.256l.221.497l.538.084A11.022 11.022 0 0 0 13 20h3v10h2v-4h3a9.01 9.01 0 0 0 9-9v-3Zm-12.898 3.947A7.51 7.51 0 0 0 5.528 14A8.895 8.895 0 0 1 4 9V4h3a8.862 8.862 0 0 1 3.026.53A4.246 4.246 0 0 0 10 5a4.966 4.966 0 0 0 5.47 4.974A8.88 8.88 0 0 1 16 13v5h-3a8.115 8.115 0 0 1-.898-.053ZM28 17a7.008 7.008 0 0 1-7 7h-3v-1a7.01 7.01 0 0 1 5.021-6.71A1.497 1.497 0 1 0 26 16.5a1.485 1.485 0 0 0-.092-.5H28Z"/>`),
		g.Group(children),
	)
}