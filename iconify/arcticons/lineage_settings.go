package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.66 18.032l1.556 3.394a.727.727 0 0 0 .358.358l3.394 1.556a.727.727 0 0 1 0 1.321l-3.394 1.556a.727.727 0 0 0-.358.358l-1.556 3.394a.727.727 0 0 1-1.32 0l-1.556-3.395a.727.727 0 0 0-.358-.358l-3.394-1.555a.727.727 0 0 1 0-1.321l3.394-1.556a.727.727 0 0 0 .358-.358l1.555-3.394a.727.727 0 0 1 1.322 0Z"/><circle cx="24" cy="24" r="14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.538 29.043h0a5.823 5.823 0 0 1 0-10.086h0a2.058 2.058 0 0 0 .753-2.812l-3.343-5.79a2.058 2.058 0 0 0-2.812-.753h0a5.823 5.823 0 0 1-8.735-5.043v0A2.058 2.058 0 0 0 27.343 2.5h-6.686a2.058 2.058 0 0 0-2.058 2.058v0a5.823 5.823 0 0 1-8.735 5.043h0a2.058 2.058 0 0 0-2.812.754l-3.343 5.79a2.058 2.058 0 0 0 .753 2.812h0a5.823 5.823 0 0 1 0 10.086h0a2.058 2.058 0 0 0-.753 2.811l3.343 5.79a2.058 2.058 0 0 0 2.812.754h0a5.823 5.823 0 0 1 8.735 5.043h0a2.058 2.058 0 0 0 2.058 2.059h6.686a2.058 2.058 0 0 0 2.058-2.058h0a5.823 5.823 0 0 1 8.735-5.043h0a2.058 2.058 0 0 0 2.812-.754l3.343-5.79a2.058 2.058 0 0 0-.753-2.812Z"/>`),
		g.Group(children),
	)
}