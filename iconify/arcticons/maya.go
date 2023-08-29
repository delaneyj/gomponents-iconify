package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.722 25.227v2.982a2.21 2.21 0 0 1-2.209 2.21h0c-.61 0-1.162-.248-1.562-.647"/><path d="M31.722 21.581v3.646c0 1.22-.989 2.209-2.209 2.209h0a2.21 2.21 0 0 1-2.21-2.21v-3.645M9.5 23.79c0-1.22.99-2.209 2.21-2.209h0c1.22 0 2.209.99 2.209 2.21v3.645M9.5 21.581v5.855m4.419-3.646c0-1.22.989-2.209 2.209-2.209h0c1.22 0 2.21.99 2.21 2.21v3.645m6.695-2.209a2.21 2.21 0 0 1-2.21 2.209h0a2.21 2.21 0 0 1-2.209-2.21v-1.435c0-1.22.99-2.21 2.21-2.21h0c1.22 0 2.209.99 2.209 2.21m0 3.645v-5.855M38.5 25.227a2.21 2.21 0 0 1-2.21 2.209h0a2.21 2.21 0 0 1-2.209-2.21v-1.435c0-1.22.99-2.21 2.21-2.21h0c1.22 0 2.209.99 2.209 2.21m0 3.645v-5.855"/></g>`),
		g.Group(children),
	)
}