package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M341 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h298q18 0 30.5-12.5T384 469V43q0-18-12.5-30.5T341 0zm0 469H43v-42h298v42zm0-85H43V107h298v277zm0-320H43V43h298v21zM213 448q0-9-6-15q-15-13-30 0q-6 6-6 15t6 15t15 6t15-6t6-15zM64 128h43v43H64v-43zm85 0h43v43h-43v-43zm86 0h42v43h-42v-43z"/>`),
		g.Group(children),
	)
}