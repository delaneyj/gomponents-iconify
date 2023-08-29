package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Financeguru(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.87 14.42c-4-5 2.9-8.43.6.28M31.36 9a47.94 47.94 0 0 0-7.49 6.21m8.2 3.4c-1.07-1.06-2.64-2.38-3.9-3.44m-12.44 4.02C17.27 17 16.57 17.66 20 13m8 8.45a1.11 1.11 0 1 1-1.1-1.15h0a1.11 1.11 0 0 1 1.1 1.11Zm-6.06 0a1.1 1.1 0 1 1-1.1-1.11h0A1.11 1.11 0 0 1 22 21.41Zm2.28 16.12a6 6 0 0 0 3.16-6.29m-3.93 2.84c1.72-1 2.12-2 2.14-4m-13.11-2.33c3.95 4.61 9.69 2.32 11.59-3.57c1 3.59 3.85 9.33 11.45 3.7m-12.2 14.82C2.52 17 45.16 17.3 25.2 43.48c.42-2.63-.87-5.85-1.88-.74Zm-8-31c1.05-6.89 14.44-12.1 17.13 0c.48 2.15-.55 5.75-.55 12.64a78.21 78.21 0 0 0-7.4-9.67c-1.83 1.63-5.74 3.92-8 9.67c-1.14-8-1.48-10.46-1.14-12.63Z"/>`),
		g.Group(children),
	)
}