package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TallahasseePoliceDepartment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.633 19.034a26.162 26.162 0 0 0 5.222-5.67a73.062 73.062 0 0 0-18.69 2.74a73.058 73.058 0 0 0-17.2 7.812a26.163 26.163 0 0 0 7.46 1.94m4.508-10.918A96.995 96.995 0 0 1 19.753 4.5s2.262 1.409 4.945 5.573"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.508 35.444a33.932 33.932 0 0 1-7.483 8.056a108.761 108.761 0 0 1-1.727-20.101c.02-.867.056-1.726.104-2.572m12.644-2.406a109.23 109.23 0 0 1 11.99 17.35a58.632 58.632 0 0 1-16.8-5.12"/>`),
		g.Group(children),
	)
}