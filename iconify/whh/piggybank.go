package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piggybank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.698 512q0 87-34 165.5t-95 139.5l59 143q11 26-.5 45t-39.5 19h-133q-28 0-55-19t-38-45l-8-19q-68 19-136 19q-39 0-77-6l-3 6q-11 26-38 45t-55 19h-133q-28 0-39.5-19t-.5-45l45-110q-95-84-130-210h-49q-26 0-45-37.5t-19-90.5t19-90.5t45-37.5h50q21-75 64-136q-5-19-14-47t-15.5-49t-12.5-46t-7-42.5t2.5-32.5t15.5-23t33-8q51 0 100 27t84 72q84-35 180-35q155 0 281 85q32-21 71-21q53 0 90.5 37.5t37.5 90.5q0 51-35 88q35 81 35 168zm-704.5-192q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5z"/>`),
		g.Group(children),
	)
}