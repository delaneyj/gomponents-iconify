package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1005 237L557 749q-18 19-44.5 19T467 749L19 237Q0 219 0 192.5T19 147L147 19q19-19 45-19h640q26 0 45 19l128 128q19 18 19 45t-19 45zM512 704l115-131q-20 3-39 3h-92q-59 0-105-10zM288 448h96q0 32 28.5 48t67.5 16h64q57 0 76.5-13t19.5-51q0-29-31-46.5T512 384q-77 0-120-1t-96.5-7.5T208 357zM192 64L64 192l65 74l-.5-4l-.5-6q0-47 45-92.5T282 91t122-27H192zm352 0q138 0 182 70l42-70H544zm325 37l-69 91H672q0-22-23-37t-51-21t-54-6h-96q-62 0-95 14.5T320 192q0 26 60 45t132 19q73 0 125 10.5t79.5 29.5t39.5 40.5t12 47.5q0 10-2 30l194-222z"/>`),
		g.Group(children),
	)
}