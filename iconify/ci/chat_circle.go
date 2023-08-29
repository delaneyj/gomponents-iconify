package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.51 19.802a9 9 0 1 0-3.312-3.312l.003.005c.073.127.11.191.127.252c.016.057.02.108.016.168a1.058 1.058 0 0 1-.07.26l-.768 2.307l-.001.003c-.162.487-.243.73-.186.892c.05.142.163.253.304.304c.162.057.404-.023.889-.185l.006-.002l2.306-.769c.131-.044.198-.066.262-.07a.471.471 0 0 1 .167.017a1.3 1.3 0 0 1 .253.127l.004.003Z"/>`),
		g.Group(children),
	)
}