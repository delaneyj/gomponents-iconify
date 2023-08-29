package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M767 1024H64q-27 0-45.5-18.5T0 960V65q0-27 18.5-45.5T64 1h384v352q0 13 9 22.5t23 9.5h351v575q0 27-18.5 45.5T767 1024zM596 640h-84V480q0-14-9.5-23t-22.5-9H352q-13 0-22.5 9t-9.5 23v160h-87q-23 0-35 26.5t-2 37.5l89 89l85 84q17 18 42 18t42-18l179-173q14-14 1-39t-38-25zM512 0q26 0 44 18l256 257q19 19 19 46H512V0z"/>`),
		g.Group(children),
	)
}