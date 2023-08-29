package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1000 66q-92 60-150 133q47 57 44 132t-56 128q-38 38-61 90t-33.5 102.5T723 752t-29.5 98t-51.5 81q-74 73-164 88t-176-21t-163-113T26 721.5T5 545t88-163q34-33 82-53t98-29.5t100.5-20T475 246t89-61q46-46 111-54.5T795 153q24-40 69.5-77T941 22t42-20q13-4 24.5 3t15.5 20q6 22-23 41z"/>`),
		g.Group(children),
	)
}