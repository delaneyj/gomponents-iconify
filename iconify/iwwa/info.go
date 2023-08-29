package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20.108 15.754a1.11 1.11 0 0 0 1.107-1.107v-1.111c0-.606-.497-1.107-1.107-1.107s-1.109.502-1.109 1.107v1.111c0 .609.499 1.107 1.109 1.107zm-2.215 4.431h1.106v6.278c0 .609.499 1.107 1.109 1.107s1.107-.498 1.107-1.107v-7.386a1.11 1.11 0 0 0-1.107-1.107h-2.215c-.61 0-1.109.498-1.109 1.107s.499 1.108 1.109 1.108zm0 0"/><path fill="currentColor" d="M33.847 34.847H6.153a1 1 0 0 1-1-1V6.153a1 1 0 0 1 1-1h27.693a1 1 0 0 1 1 1v27.693a1 1 0 0 1-.999 1.001zm-26.694-2h25.693V7.153H7.153v25.694z"/>`),
		g.Group(children),
	)
}