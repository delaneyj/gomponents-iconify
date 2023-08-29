package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M6 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm8 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0-16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm8 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M4 6v12m8-12v12m8-12v2c0 1.886 0 2.828-.586 3.414C18.828 12 17.886 12 16 12H4"/><path fill="currentColor" d="M18 15v-.75a.75.75 0 0 0-.75.75H18Zm-.75 7a.75.75 0 0 0 1.5 0h-1.5Zm4.11.392a.75.75 0 1 0 1.28-.784l-1.28.784ZM18 15.75h2.286v-1.5H18v1.5Zm.75 2.75V15h-1.5v3.5h1.5Zm2.5-1.75c0 .567-.446 1-.964 1v1.5c1.375 0 2.464-1.134 2.464-2.5h-1.5Zm-.964-1c.518 0 .964.433.964 1h1.5c0-1.366-1.089-2.5-2.464-2.5v1.5Zm0 2h-.429v1.5h.429v-1.5Zm-.429 0H18v1.5h1.857v-1.5Zm-.64 1.142l2.143 3.5l1.28-.784l-2.143-3.5l-1.28.784ZM17.25 18.5V22h1.5v-3.5h-1.5Z"/></g>`),
		g.Group(children),
	)
}