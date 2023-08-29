package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ftpaccounts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.678 768h-320v128h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128V768h-320q-53 0-90.5-37.5T.678 640V256q0-26 19-45t45-19h480q0-8 11.5-29t22.5-38l12-17q7-19 28-31.5t44-12.5h240q24 0 45 12.5t28 31.5l49 84v448q0 53-37.5 90.5t-90.5 37.5zm-320-186v-53q64-37 64-168q0-52-36-78.5t-92-26.5t-92 26.5t-36 78.5q0 125 64 166v55q-56 11-91.5 37t-36.5 58q0 27 192 27t192-27q-1-32-36.5-58t-91.5-37zm-530-538q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84h-480z"/>`),
		g.Group(children),
	)
}