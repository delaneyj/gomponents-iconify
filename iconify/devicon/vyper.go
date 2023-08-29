package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vyper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path d="M80 91.712L64 64L48 91.712l16 27.712z" opacity=".8"/><path d="M96 64L80 36.287L64 64l16 27.712zm-32 0L48 36.287L32 64l16 27.712z" opacity=".6"/><path d="M112 36.287L96 8.574L80 36.287L96 64zm-64 0L64 64l16-27.713l-8-13.857H56zm0 0L32 8.574L16 36.287L32 64z" opacity=".45"/><path d="M128 8.574H96l16 27.712zm-96 0H0l16 27.712z" opacity=".3"/>`),
		g.Group(children),
	)
}