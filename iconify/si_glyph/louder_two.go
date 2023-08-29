package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LouderTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.962" cy="10.954" rx="1.962" ry="1.954"/><path d="M8.547.014H1.465C.675.014.031.612.031 1.351v13.295c0 .737.644 1.337 1.434 1.337h7.082c.789 0 1.433-.6 1.433-1.337V1.351C9.979.611 9.336.014 8.547.014zM4.986 1.948c1.144 0 2.066.927 2.066 2.069a2.066 2.066 0 1 1-4.131 0c0-1.142.926-2.069 2.065-2.069zm.006 12.083a3.038 3.038 0 0 1-3.047-3.033a3.038 3.038 0 0 1 3.047-3.031a3.039 3.039 0 0 1 3.047 3.031a3.04 3.04 0 0 1-3.047 3.033z"/><ellipse cx="4.991" cy="3.978" rx=".991" ry=".978"/></g>`),
		g.Group(children),
	)
}