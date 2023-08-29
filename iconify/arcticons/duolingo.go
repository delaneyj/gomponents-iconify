package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duolingo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="11.188" height="14.476" x="6.999" y="18.918" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.594"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.06 22.11a2.538 2.538 0 0 0-.959.188a2.124 2.124 0 0 1-1.588 3.166v2.288a2.547 2.547 0 0 0 5.094 0v-3.095a2.547 2.547 0 0 0-2.547-2.547Z"/><rect width="11.188" height="14.476" x="29.813" y="18.918" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.594"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.865 22.11a2.538 2.538 0 0 0-.959.188a2.124 2.124 0 0 1-1.588 3.166v2.288a2.547 2.547 0 1 0 5.094 0v-3.095a2.547 2.547 0 0 0-2.547-2.547Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.077 33.256a8.824 8.824 0 0 0 7.595 4.324h0a8.828 8.828 0 0 0 8.828-8.828v-5.19a8.794 8.794 0 0 0-2.21-5.844h0c0-1.926-.468-4.617-1.771-4.617s-2.509 1.945-2.509 1.945s.815-4.626-1.524-4.626c-3.327 0-5.084 8.956-12.485 8.956s-9.16-8.956-12.486-8.956c-2.34 0-1.525 4.626-1.525 4.626s-1.205-1.945-2.509-1.945s-1.77 2.691-1.77 4.617h0A8.794 8.794 0 0 0 3.5 23.562v5.19a8.828 8.828 0 0 0 8.828 8.828h0a8.824 8.824 0 0 0 7.595-4.324"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 33.833l4.9-.694a4.504 4.504 0 0 0-4.9-4.222a4.504 4.504 0 0 0-4.9 4.222Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.578 33.326a3.582 3.582 0 1 1-7.156 0"/>`),
		g.Group(children),
	)
}