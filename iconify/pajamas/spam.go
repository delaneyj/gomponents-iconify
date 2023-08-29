package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.103 1.529L8 0L5.897 1.529l-2.6-.001L2.496 4L.392 5.528L1.196 8l-.804 2.472L2.495 12l.803 2.473l2.6-.001L8 16l2.103-1.529l2.6.001l.802-2.473l2.103-1.527L14.804 8l.804-2.472L13.505 4l-.803-2.473l-2.6.001Zm1.51 1.5H9.614l-.395-.287L8 1.855l-1.22.887l-.395.287H4.387l-.465 1.435l-.15.464l-.395.287l-1.222.886l.467 1.435l.151.464l-.15.464L2.154 9.9l1.222.886l.395.287l.15.464l.466 1.436l1.509-.001h.488l.395.287l1.22.887l1.22-.887l.395-.287h1.998l.465-1.435l.15-.464l.395-.287l1.222-.886l-.467-1.435l-.15-.465l.15-.464l.467-1.435l-1.22-.886l-.396-.287l-.15-.464l-.466-1.436ZM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0v-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}