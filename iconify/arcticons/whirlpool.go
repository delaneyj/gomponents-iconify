package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whirlpool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.38 26.313a1.651 1.651 0 0 1-1.646-1.646v-1.07a1.646 1.646 0 1 1 3.292 0v1.07a1.651 1.651 0 0 1-1.646 1.646Zm4.82-.002h0a1.651 1.651 0 0 1-1.646-1.646v-1.07A1.651 1.651 0 0 1 37.2 21.95h0a1.651 1.651 0 0 1 1.646 1.646v1.07a1.651 1.651 0 0 1-1.646 1.646Zm-15.904-2.844a1.651 1.651 0 0 1 1.647-1.647m-1.647 0v4.363"/><circle cx="19.394" cy="19.861" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.394 21.837V26.2m4.657-6.563V25.4a.778.778 0 0 0 .824.824h.247m15.308-6.625v5.762a.778.778 0 0 0 .823.823h.247m-15.368-1.567a1.646 1.646 0 0 0 3.292 0v-1.07a1.646 1.646 0 0 0-3.292 0m0-1.647v6.585m-13.047-8.89l-1.646 6.585l-1.646-6.585l-1.647 6.585L6.5 19.595m7.906.05v6.585m0-2.716a1.646 1.646 0 0 1 3.293 0v2.716"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.333 20.843c2.895-1.867 5.368-2.74 5.805-2.452c3.244 1.823-9.393 13.164-11.825 10.998c-.565-.407-.174-1.891.976-3.705"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}