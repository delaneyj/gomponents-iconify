package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bandage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.927 10.543C-.238 11.711-.25 13.59.902 14.742l.354.354c1.151 1.15 3.029 1.139 4.195-.027l1.536-1.534l-4.525-4.523l-1.535 1.531zm14.202-9.321l-.354-.354C13.623-.283 11.742-.27 10.57.899L9.008 2.463l4.527 4.525l1.562-1.563c1.169-1.169 1.182-3.052.032-4.203zM3.469 7.969L8 12.469l4.562-4.438l-4.588-4.646l-4.505 4.584zm4.557-1.943H6.979V4.979h1.047v1.047zm2 2.016H8.979V6.98h1.047v1.062zm-.015 2H8.98V8.98h1.031v1.062zm-2-2H6.98V6.98h1.031v1.062zm-2 0H4.98V6.98h1.031v1.062zm2.015 2H6.979V8.98h1.047v1.062z"/>`),
		g.Group(children),
	)
}