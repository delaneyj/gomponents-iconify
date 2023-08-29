package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedRealityBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Zm7.216-4.224a.75.75 0 0 0-1.432 0l-2.5 8a.75.75 0 0 0 1.432.448l.71-2.274h2.148l.71 2.274a.75.75 0 0 0 1.432-.448l-2.5-8Zm-1.32 4.674h1.209L8.5 10.515l-.605 1.935ZM13.25 8a.75.75 0 0 1 .75-.75h2a2.75 2.75 0 0 1 .783 5.387l1.853 2.965a.75.75 0 1 1-1.272.796l-2.28-3.648h-.334V16a.75.75 0 0 1-1.5 0V8Zm1.5 3.25v-2.5H16a1.25 1.25 0 1 1 0 2.5h-1.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}