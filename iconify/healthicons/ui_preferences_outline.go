package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiPreferencesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M23 22.922c0-.526.474-.922 1-.922s1 .396 1 .922V41a1 1 0 1 1-2 0V22.922Z"/><path fill-rule="evenodd" d="M20 16a4.002 4.002 0 0 1 3-3.874V7a1 1 0 1 1 2 0v5.126A4.002 4.002 0 0 1 24 20a4 4 0 0 1-4-4Zm4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-8 14a4 4 0 1 1-4.992-3.876A1.017 1.017 0 0 1 11 28V7a1 1 0 1 1 2 0v21c0 .042-.003.083-.008.124A4.002 4.002 0 0 1 16 32Zm-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/><path d="M36 42a1 1 0 0 1-1-1v-7.078c0-.526.474-.922 1-.922s1 .396 1 .922V41a1 1 0 0 1-1 1Z"/><path fill-rule="evenodd" d="M36 6a1 1 0 0 0-1 1v16.126A4.002 4.002 0 0 0 36 31a4 4 0 0 0 1-7.874V7a1 1 0 0 0-1-1Zm0 19a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z" clip-rule="evenodd"/><path d="M11 41a1 1 0 1 0 2 0v-2.078c0-.526-.474-.922-1-.922s-1 .396-1 .922V41Z"/></g>`),
		g.Group(children),
	)
}