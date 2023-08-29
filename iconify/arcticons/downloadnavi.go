package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downloadnavi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.35 24h-5.51v-9.58h0a.82.82 0 0 0-.82-.83H21a.83.83 0 0 0-.81.84V24h-5.48a.9.9 0 0 0-.78.32a.91.91 0 0 0 .07 1.27l9.3 9.87h0a.93.93 0 0 0 1.31 0L34 25.59v-.06a.86.86 0 0 0 .25-.6a.89.89 0 0 0-.9-.93Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m45.29 23.68l-2.16-1.2a19.11 19.11 0 0 0-4.79-11.11l.54-2a.22.22 0 0 0 0-.16a.27.27 0 0 0-.33-.17l-2 .53a19.1 19.1 0 0 0-11.13-4.68l-1.2-2.15a.36.36 0 0 0-.63 0l-1.2 2.16a19.11 19.11 0 0 0-11.11 4.79l-2-.54a.22.22 0 0 0-.16 0a.27.27 0 0 0-.12.33l.53 2A19.1 19.1 0 0 0 4.8 22.61l-2.15 1.2a.36.36 0 0 0 0 .63l2.16 1.2A19.11 19.11 0 0 0 9.6 36.75l-.54 2a.22.22 0 0 0 0 .16a.27.27 0 0 0 .33.17l2-.53a19.1 19.1 0 0 0 11.13 4.69l1.2 2.15a.36.36 0 0 0 .63 0l1.2-2.16a19.11 19.11 0 0 0 11.11-4.79l2 .54a.22.22 0 0 0 .16 0a.27.27 0 0 0 .17-.33l-.53-2a19.1 19.1 0 0 0 4.69-11.13l2.15-1.2a.36.36 0 0 0 0-.63ZM24 40.32a16.34 16.34 0 1 1 0-32.67h0a16.34 16.34 0 0 1 0 32.67Z"/>`),
		g.Group(children),
	)
}