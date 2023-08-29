package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.936 512h-896q-27 0-45.5-18.5T.936 448V320q0-27 18.5-45.5t45.5-18.5h896q26 0 45 18.5t19 45.5v128q0 26-19 45t-45 19zm-640-320h-256q-27 0-45.5-18.5T.936 128V96q0-50 35-73t93-23q38 0 97 16.5t109 47.5t50 64q0 27-18.5 45.5t-45.5 18.5zm-64 384v128q0 27-19 45.5t-45.5 18.5t-45-18.5t-18.5-45.5V576h128zm640 0v128q0 27-19 45.5t-45.5 18.5t-45-18.5t-18.5-45.5V576h128z"/>`),
		g.Group(children),
	)
}