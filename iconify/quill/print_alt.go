package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h-3.621c-.252 0-.504-.022-.752-.067l-.92-.164c-3.383-.604-5.244-3.076-5.62-5.769c-.484-3.466 1.494-7.3 5.861-7.992a.114.114 0 0 0 .08-.053l.636-1.049C13.44 3.982 16.582 3.56 19 4.687c2.378 1.11 4.055 3.722 3.037 6.93a.11.11 0 0 0 .084.142c2.952.586 4.372 3.143 4.173 5.567a5.188 5.188 0 0 1-.915 2.566c-1.353 1.92-4.028 2.1-6.377 2.115c-1 .006-2-.007-3.002-.007Z"/><path d="m14.5 22l-.5 6h4l-.5-6h-3Z"/></g>`),
		g.Group(children),
	)
}