package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvoiceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 9.75a.75.75 0 0 0-.75-.75h-6a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75Zm-1 3a.75.75 0 0 0-.75-.75h-5a.75.75 0 1 0 0 1.5h5a.75.75 0 0 0 .75-.75Zm.25 2.25a.75.75 0 1 1 0 1.5h-6a.75.75 0 0 1 0-1.5h6Z"/><path fill="currentColor" fill-rule="evenodd" d="M6 21.75h13A2.75 2.75 0 0 0 21.75 19v-5.5a.75.75 0 0 0-.75-.75h-3.25V4.943c0-1.423-1.609-2.251-2.767-1.424l-.175.125a2.26 2.26 0 0 1-2.622-.004a3.77 3.77 0 0 0-4.372 0a2.26 2.26 0 0 1-2.622.004l-.175-.125c-1.158-.827-2.767 0-2.767 1.424V18A3.75 3.75 0 0 0 6 21.75ZM8.686 4.86a2.27 2.27 0 0 1 2.628 0a3.76 3.76 0 0 0 4.366.005l.175-.125a.25.25 0 0 1 .395.203V19c0 .45.108.875.3 1.25H6A2.25 2.25 0 0 1 3.75 18V4.943a.25.25 0 0 1 .395-.203l.175.125a3.76 3.76 0 0 0 4.366-.005ZM17.75 19v-4.75h2.5V19a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}