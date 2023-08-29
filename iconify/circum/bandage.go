package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bandage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.435 7.5H6.565a4.5 4.5 0 0 0 0 9h10.87a4.5 4.5 0 0 0 0-9Zm-9.93 8h-.94a3.5 3.5 0 0 1 0-7h.94Zm8 0h-7v-7h7Zm1.93 0h-.93v-7h.93a3.5 3.5 0 0 1 0 7Z"/><circle cx="10.252" cy="10.501" r=".625" fill="currentColor"/><circle cx="10.252" cy="13.501" r=".625" fill="currentColor"/><circle cx="13.752" cy="10.5" r=".625" fill="currentColor"/><circle cx="13.752" cy="13.5" r=".625" fill="currentColor"/>`),
		g.Group(children),
	)
}