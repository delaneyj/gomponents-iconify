package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 759q-97 62-96 137q1 36 4 55q92 33 92 73H0q0-40 93-74q3-28 3-54q-2-77-96-137q255-11 137-237q227 117 238-138q60 94 137 96q75 2 137-96q11 255 238 138q-118 226 137 237zM512 576q-87 0-161 43T234.5 735.5T192 896q0 12 1 28q30-6 64-11q-1-15-1-17q0-106 75-181t181-75t181 75t75 181q0 2-1 17q34 5 64 11q1-17 1-28q0-87-43-160.5T672.5 619T512 576zm24-265q-10 9-24 9t-23-9L327 165q-17-15 8-37h113V32q0-13 9.5-22.5T480 0h64q13 0 22.5 9.5T576 32v96h113q25 22 9 37z"/>`),
		g.Group(children),
	)
}