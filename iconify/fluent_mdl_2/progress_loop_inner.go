package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgressLoopInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M448 1024q0 60 12 118t36 112l-111 64q-32-69-48-143t-17-151q0-91 22-176t64-160t99-138t129-111t153-78t173-38v129q-109 12-202 61T595 640T487 815t-39 209zm640-701q90 8 172 38t154 77t129 111t99 139t63 160t23 176q0 77-16 151t-49 144l-111-64q23-55 35-113t13-118q0-110-39-208t-108-176t-162-126t-203-62V323zm-64 1277q69 0 134-16t125-46t111-74t93-99l111 65q-50 70-113 125t-137 94t-156 58t-168 21q-86 0-168-20t-156-59t-138-94t-113-126l112-64q41 55 92 98t111 74t125 47t135 16z"/>`),
		g.Group(children),
	)
}