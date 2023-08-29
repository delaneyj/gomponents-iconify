package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCircleAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v-3m0 0V9m0 3H9m3 0h3m-3 9a8.96 8.96 0 0 1-4.49-1.198a1.276 1.276 0 0 0-.257-.13a.475.475 0 0 0-.167-.017a1.12 1.12 0 0 0-.258.07l-2.31.769h-.002c-.487.163-.731.245-.893.187a.5.5 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306l.002-.005c.042-.129.064-.194.068-.256a.478.478 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21Z"/>`),
		g.Group(children),
	)
}