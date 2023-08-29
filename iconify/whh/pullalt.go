package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pullalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.59 512q-139 0-257-17T69.094 448.5T.59 384V256q0-8 6-18q62 36 202.5 59t303.5 23t303.5-23t202.5-59q6 10 6 18v128q0 35-68.5 64.5T769.59 495t-257 17zm0-256q-139 0-257-17T69.094 192.5T.59 128t68.5-64.5T255.59 17t257-17t257 17t186.5 46.5t68.5 64.5t-68.5 64.5t-186.5 46.5t-257 17zm-142 527l78.002 62V640q0-26 18.5-45t45.5-19t45.5 19t18.5 45v204l77-61q20-16 47.5-16t47 16t19.5 38t-19 37l-189 150q-20 16-47.5 16t-47.5-16l-189-150q-19-15-19-37t19.5-38t47-16t47.5 16z"/>`),
		g.Group(children),
	)
}