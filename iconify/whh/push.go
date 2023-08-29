package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Push(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 1024H64q-27 0-45.5-18.5T0 960V65q0-27 18.5-45.5T64 1h384v352q0 13 9 22.5t23 9.5h351v575q0 27-18.5 45.5T768 1024zM633 640L454 466q-17-18-42-18t-43 18l-85 85l-88 89q-11 10 1.5 37t34.5 27h88v159q0 14 9 23t23 9h128q13 0 22.5-9t9.5-23V704h84q25 0 38-25.5t-1-38.5zM512 0q25 0 44 18l256 257q19 19 19 46H512V0z"/>`),
		g.Group(children),
	)
}