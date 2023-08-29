package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtoMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.057" cy="32.586" r="2.164" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.613" cy="18.736" r="2.164" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.22 32.586V15.414m.002 6.472a11.153 11.153 0 0 0-4.096-.05a10.918 10.918 0 0 0 2.508 21.65a11.174 11.174 0 0 0 10.337-11.29v-14a4.946 4.946 0 0 1 4.946-4.945h2.94V4.5h-5.722A10.914 10.914 0 0 0 22.22 15.414"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.135 4.5H13.169v8.751h9.051"/>`),
		g.Group(children),
	)
}