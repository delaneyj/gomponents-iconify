package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkFilledWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 2a12 12 0 1 0 3.393 23.506l3.924-7.85A2.985 2.985 0 0 1 24 16h1.82A11.935 11.935 0 0 0 14 2Zm-2 16.59l-4-4L9.59 13L12 15.41L17.41 10L19 11.59Z"/><path fill="currentColor" d="M27.38 28h-6.762L24 21.236ZM24 18a1 1 0 0 0-.895.553l-5 10A1 1 0 0 0 19 30h10a1 1 0 0 0 .921-1.39l-5.026-10.057A1 1 0 0 0 24 18Z"/><path fill="none" d="m12 18.591l-4-4L9.591 13L12 15.409L17.409 10L19 11.591l-7 7z"/>`),
		g.Group(children),
	)
}