package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onecard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 10h-33a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h33a3 3 0 0 0 3-3V13a3 3 0 0 0-3-3Zm-36 22h39"/><rect width="7.614" height="10.088" x="8.697" y="16" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.807" ry="3.807"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.783 26.088v-6.281A3.807 3.807 0 0 0 23.976 16h0a3.807 3.807 0 0 0-3.807 3.807m0 6.281V16m18.635 8.167a3.806 3.806 0 0 1-3.308 1.921h0a3.807 3.807 0 0 1-3.807-3.807v-2.474A3.807 3.807 0 0 1 35.496 16h0a3.807 3.807 0 0 1 3.807 3.807v1.237h-7.614M12.504 18.39v5.309m-1.12-4.586h.493a.627.627 0 0 0 .627-.627"/>`),
		g.Group(children),
	)
}