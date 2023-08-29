package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.206 37.317l6.678 2.421V15.005a.996.996 0 0 1 1.336-.931h0a2.022 2.022 0 0 1 1.335 1.9V26.18a5.032 5.032 0 0 0 6.678-4.843a10.376 10.376 0 0 0-6.853-9.748l-9.174-3.326Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.53 26.584l-10.28 3.68a2.645 2.645 0 0 0-1.589 3.394a4.005 4.005 0 0 0 1.813 1.815a14.164 14.164 0 0 0 10.658.518c.099-.03 1.976-.647 2.074-.68v-3.873l-5.785 2.038a4.26 4.26 0 0 1-2.788.031c-.269-.097-.695-.628.117-1l5.78-2.048Zm9.354 2.528c.254-.087 2.228-.77 2.484-.846a20.677 20.677 0 0 1 13.325.51a2.832 2.832 0 0 1 1.61 3.676a2.91 2.91 0 0 1-1.744 1.652l-15.675 5.635v-3.875l11.726-4.202a.776.776 0 0 0 0-1.46a5.878 5.878 0 0 0-4.007.007l-7.719 2.777Z"/>`),
		g.Group(children),
	)
}