package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.394 24a7.877 7.877 0 0 0-1.67-8.568a3.918 3.918 0 0 0-1.085-4.414l2.776-2.776a2.002 2.002 0 0 0 0-2.829l-2.828-2.828a2.002 2.002 0 0 0-2.829 0L6.586 15.757a2.003 2.003 0 0 0 0 2.829l2.828 2.828a2.002 2.002 0 0 0 2.829 0l4.775-4.775a3.933 3.933 0 0 0 5.513.432A5.944 5.944 0 0 1 23.177 24H16v4H4v2h24v-6Zm-14.566-4L8 17.171l1.879-1.878l2.828 2.828ZM16 14a3.981 3.981 0 0 0 .076.752l-1.955 1.955l-2.828-2.828l9.88-9.88L24 6.827l-3.249 3.25A3.977 3.977 0 0 0 16 14Zm4 2a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm6 12h-8v-2h8Z"/>`),
		g.Group(children),
	)
}