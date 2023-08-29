package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NiagaraLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.323 35.437a18.848 18.848 0 0 0 9.438 7.682a1.184 1.184 0 0 0 1.192-.28c11.293-12.185 7.68-26.43-2.213-32.423c-1.035-.627-1.642-.09-1.361 1.086c2.35 9.85-.519 16.95-6.855 22.651a1.047 1.047 0 0 0-.2 1.284Z"/><path fill="none" stroke="currentColor" d="M36.136 34.721c-3.332-8.545-3.208-16.795 3.183-24.513a.911.911 0 0 0-.012-1.167a17.219 17.219 0 0 0-8.094-4.788a2.329 2.329 0 0 0-2.19.585c-2.627 2.96-3.14 5.99-4.25 9a.871.871 0 0 0 .039.616c2.252 4.356 4.168 8.945 3.67 15.203a1.141 1.141 0 0 0 .206.696a20.214 20.214 0 0 0 6.456 5.293c.947.512 1.383.078.992-.925Z"/>`),
		g.Group(children),
	)
}