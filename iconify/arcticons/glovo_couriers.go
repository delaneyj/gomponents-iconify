package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlovoCouriers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.983 4.5a12.235 12.235 0 0 0-12.227 12.252a12.046 12.046 0 0 0 2.307 7.133l.33.462l6.36 9.016a2.998 2.998 0 0 0 2.472 1.287h1.55a2.996 2.996 0 0 0 2.472-1.287l6.36-9.016l.33-.462a12.114 12.114 0 0 0 2.307-7.132A12.285 12.285 0 0 0 23.983 4.5m5.042 15.884l-.33.462l-4.68 6.638l-4.712-6.638l-.33-.495a6.289 6.289 0 0 1-1.186-3.632a6.195 6.195 0 0 1 6.182-6.209h.014a6.195 6.195 0 0 1 6.196 6.194v.014a6.36 6.36 0 0 1-1.154 3.665m-8.47 19.781a3.345 3.345 0 0 1 3.395-3.368a3.298 3.298 0 0 1 3.395 3.335v.034a3.395 3.395 0 0 1-6.79 0"/>`),
		g.Group(children),
	)
}