package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.998 760v168q0 13-9.5 22.5t-22.5 9.5h-168q-116 64-248 64t-248-64h-168q-13 0-22.5-9.5t-9.5-22.5V760q-64-116-64-248t64-248V96q0-13 9.5-22.5t22.5-9.5h168q116-64 248-64t248 64h168q13 0 22.5 9.5t9.5 22.5v168q64 116 64 248t-64 248zm-448-632q-84 0-158 34t-128 94h158l-51 51l75 303l104-290l103 290l76-303l-51-51h158q-54-60-128-94t-158-34zm315 165l-6 6l-106 426l21 43h-192l19-38l-51-154l-51 154l19 38h-192l21-43l-106-426l-6-6q-69 100-69 219q0 104 51.5 192.5t140 140t192.5 51.5t192.5-51.5t140-140t51.5-192.5q0-119-69-219z"/>`),
		g.Group(children),
	)
}